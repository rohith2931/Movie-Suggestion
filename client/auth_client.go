package client

import (
	"context"
	pb "example/movieSuggestion/msproto"
	"time"

	"google.golang.org/grpc"
)

// AuthClient is used validate the user
type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

// This function returns New AuthClient
func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

// This RPC is used to send jwt token to the user/client
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}
	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil
}
