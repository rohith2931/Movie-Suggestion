package server

import (
	"context"
	"log"

	"example/movieSuggestion/authorization"
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	jwtManager *authorization.JWTManager
	Db         *gorm.DB
}

// NewAuthServer returns a new auth server
func NewAuthServer(jwtManager *authorization.JWTManager, db *gorm.DB) pb.AuthServiceServer {
	return &AuthServer{jwtManager: jwtManager, Db: db}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("Server : Logging in the User")
	user := schema.User{
		UserName: req.GetUsername(),
		Password: req.GetPassword(),
	}
	var count int64
	server.Db.Model(&schema.User{}).Where("user_name=? and password=?", user.UserName, user.Password).Find(&user).Count(&count)
	if count != 0 {
		token, err := server.jwtManager.Generate(&user)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate access token")
		}

		res := &pb.LoginResponse{AccessToken: token}
		return res, nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "Not a valid user")
}
