package authorization

import (
	"context"

	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/schema"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is used to authenticate the user
type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	jwtManager *JWTManager
	Db         *gorm.DB
}

// NewAuthServer returns a new auth server
func NewAuthServer(jwtManager *JWTManager, db *gorm.DB) pb.AuthServiceServer {
	return &AuthServer{jwtManager: jwtManager, Db: db}
}

// This RPC is used to login the user and return jwt token as a response
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
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
