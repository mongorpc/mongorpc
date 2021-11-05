package mongorpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/golang-jwt/jwt"
	"github.com/mongorpc/mongorpc/proto"
	"github.com/osohq/go-oso"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	JWTSecret   string
	Oso         oso.Oso
	PolicyFiles []string
}

// RouteInfoPayload is the payload of the route info
type RouteInfoPayload struct {
	Database   string  `json:"database"`
	Collection *string `json:"collection"`
	DocumentID *string `json:"document_id"`
	UID        string  `json:"uid"`
}

// Auth interceptor for validating JWT token
func (interceptor *Interceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Infoln("method: ", info.FullMethod)
	err := interceptor.authorize(ctx, info.FullMethod, req)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

// Auth interceptor for validating JWT token for streams
func (interceptor *Interceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Infoln("method: ", info.FullMethod)
	err := interceptor.authorize(stream.Context(), info.FullMethod, srv)
	if err != nil {
		return err
	}
	return handler(srv, stream)
}

// Authorize validates JWT token
func (interceptor *Interceptor) authorize(ctx context.Context, method string, payload interface{}) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]

	token, err := jwt.Parse(accessToken, interceptor.keyFunc)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for _, m := range proto.MongoRPC_ServiceDesc.Methods {
			fullMethod := fmt.Sprintf("/%s/%s", proto.MongoRPC_ServiceDesc.ServiceName, m.MethodName)

			if fullMethod == method {
				return interceptor.CheckPermission(method, claims, payload)
			}
		}
		return status.Error(codes.PermissionDenied, "permission denied")
	} else {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

}

// CheckPermission checks if the user has permission to access the route
func (interceptor *Interceptor) CheckPermission(method string, claims jwt.MapClaims, payload interface{}) error {

	v := RouteInfoPayload{
		UID: claims["uid"].(string),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return status.Errorf(codes.PermissionDenied, "unable to parse payload")
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return status.Errorf(codes.PermissionDenied, "invalid payload")
	}

	if v.UID == "" {
		return status.Errorf(codes.PermissionDenied, "invalid uid")
	}

	if v.Database == "" {
		return status.Errorf(codes.PermissionDenied, "invalid database")
	}

	route := fmt.Sprintf("/%s", v.Database)
	if v.Collection != nil {
		route = fmt.Sprintf("%s/%s", route, *v.Collection)
	}

	if v.DocumentID != nil {
		route = fmt.Sprintf("%s/%s", route, *v.DocumentID)
	}

	action := "read"

	if method == "/mongorpc.MongoRPC/CreateDocument" ||
		method == "/mongorpc.MongoRPC/UpdateDocument" ||
		method == "/mongorpc.MongoRPC/DeleteDocument" ||
		method == "/mongorpc.MongoRPC/CreateIndex" ||
		method == "/mongorpc.MongoRPC/ListIndexes" ||
		method == "/mongorpc.MongoRPC/DeleteIndex" ||
		method == "/mongorpc.MongoRPC/Reindex" ||
		method == "/mongorpc.MongoRPC/CreateCollection" ||
		method == "/mongorpc.MongoRPC/RenameCollection" ||
		method == "/mongorpc.MongoRPC/DeleteCollection" {
		action = "write"
	}

	type AuthorizePayload struct {
		Actor    jwt.MapClaims `json:"actor"`
		Action   string        `json:"action"`
		Resource string        `json:"resource"`
	}

	logrus.Println("Checking permission for:", AuthorizePayload{
		Actor:    claims,
		Action:   action,
		Resource: route,
	})

	err = interceptor.Oso.Authorize(claims, action, route)
	if err != nil {
		return status.Errorf(codes.PermissionDenied, "permission denied")
	}

	return nil
}

// KeyFunc is a callback function to generate key based on the JWT token
func (interceptor *Interceptor) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(interceptor.JWTSecret), nil
}

func (interceptor *Interceptor) LoadOSOPolicy() error {
	err := interceptor.Oso.ClearRules()
	if err != nil {
		return err
	}
	err = interceptor.Oso.LoadFiles(interceptor.PolicyFiles)
	if err != nil {
		return err
	}

	logrus.Println("Loaded Oso policy files:", interceptor.PolicyFiles)
	return nil
}

func (interceptor *Interceptor) WatchFile(err error, event *fsnotify.Event) {
	if err != nil {
		logrus.Errorf("error watching file: %v", err)
		return
	}
	err = interceptor.LoadOSOPolicy()
	if err != nil {
		logrus.Errorf("Failed to reload policy: %v", err)
	}
}
