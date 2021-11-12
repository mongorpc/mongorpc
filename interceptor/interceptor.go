package interceptor

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Interceptor struct {
	// Oso       *oso.Oso
	JWTSecret string
}

func (i *Interceptor) UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logrus.Infoln("method: ", info.FullMethod)

	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	// }
	// logrus.Println(md)

	logrus.Println("JWT Secret: ", i)
	// if i.JWTSecret != nil {
	err := i.Authorize(ctx)
	if err != nil {
		return nil, err
	}
	// }
	return handler(ctx, req)
}

func (i *Interceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Infoln("method: ", info.FullMethod)

	// if i.JWTSecret != nil {
	err := i.Authorize(stream.Context())
	if err != nil {
		return err
	}
	// }
	return handler(srv, stream)
}

func (interceptor *Interceptor) Authorize(ctx context.Context) error {

	claims, err := interceptor.ValidateJWT(ctx)
	if err != nil {
		return err
	}
	logrus.Debug("claims: ", claims)

	return nil
}

// Authorize validates JWT token
func (interceptor *Interceptor) ValidateJWT(ctx context.Context) (*jwt.Claims, error) {

	token, err := interceptor.ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	return &token.Claims, nil
}

func (interceptor *Interceptor) ExtractToken(ctx context.Context) (*jwt.Token, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]

	accessToken, err := stripBearerPrefixFromTokenString(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	token, err := jwt.Parse(accessToken, interceptor.keyFunc)

	return token, err
}

// KeyFunc is a callback function to generate key based on the JWT token
func (interceptor *Interceptor) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(interceptor.JWTSecret), nil
}

// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}
