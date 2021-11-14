package interceptors

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JWTInterceptor struct {
	secret string
	claims interface{}
}

func NewJWTInterceptor(secret string, claims interface{}) *JWTInterceptor {
	return &JWTInterceptor{
		secret: secret,
		claims: claims,
	}
}

func (interceptor *JWTInterceptor) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// extract and validate token
	token, err := interceptor.ExtractToken(ctx)
	if err != nil {
		return nil, err
	}

	// create new context
	newContext := context.WithValue(ctx, "token", token)

	// return new context and request
	return handler(newContext, req)
}

func (interceptor *JWTInterceptor) StreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// create new wrapper for modifying the context
	wrapper := NewStreamContextWrapper(stream)

	token, err := interceptor.ExtractToken(stream.Context())
	if err != nil {
		return err
	}
	ctx := context.WithValue(stream.Context(), "token", token)
	wrapper.SetContext(ctx)

	// return handler(srv, wrapper)
	return handler(srv, wrapper)
}

// Extracts and Validates token from metadata
func (interceptor *JWTInterceptor) ExtractToken(ctx context.Context) (*jwt.Token, error) {
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

	token, err := jwt.ParseWithClaims(accessToken, &interceptor.claims, interceptor.keyFunc)

	return token, err
}

// KeyFunc is a callback function to generate key based on the JWT token
func (interceptor *JWTInterceptor) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(interceptor.secret), nil
}

// Strips 'Bearer ' prefix from bearer token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}
