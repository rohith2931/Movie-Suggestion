package authorization

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor is a server interceptor used  for authentication and authorization
type ServerAuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string][]string
}

// NewAuthInterceptor returns a new auth interceptor
func NewServerAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string][]string) *ServerAuthInterceptor {
	return &ServerAuthInterceptor{jwtManager, accessibleRoles}
}

// This function returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *ServerAuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}

// This function returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *ServerAuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// log.Println("--> stream interceptor: ", info.FullMethod)
		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, stream)
	}
}

// This function is used to check the user is authorized to use the endpoint/RPC
func (interceptor *ServerAuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		return nil
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}
	for _, role := range accessibleRoles {
		fmt.Println(claims)
		if role == claims.Role {
			return nil
		}
	}
	return status.Error(codes.PermissionDenied, "No permission to access this RPC")
}
