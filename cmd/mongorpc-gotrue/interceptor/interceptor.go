package interceptor

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	gotrue "github.com/netlify/gotrue/api"
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

	err := i.Authorize(ctx, req)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func (i *Interceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logrus.Infoln("method: ", info.FullMethod)

	err := i.Authorize(stream.Context(), srv)
	if err != nil {
		return err
	}
	return handler(srv, stream)
}

func (interceptor *Interceptor) Authorize(ctx context.Context, req interface{}) error {

	token, err := interceptor.ExtractToken(ctx)
	if err != nil {
		return err
	}

	claims := token.Claims.(*gotrue.GoTrueClaims)
	err = interceptor.Permissions(ctx, claims, req)
	if err != nil {
		return err
	}

	return nil
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

	token, err := jwt.ParseWithClaims(accessToken, &gotrue.GoTrueClaims{}, interceptor.keyFunc)

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
