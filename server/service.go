package server

import (
	pb "example/movieSuggestion/msproto"

	"github.com/jinzhu/gorm"
)

type MsServer struct {
	pb.UnimplementedMsDatabaseServer
	Db *gorm.DB
}
