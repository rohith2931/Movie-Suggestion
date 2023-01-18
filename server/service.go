package server

import (
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
)

type MsServer struct {
	pb.UnimplementedMsDatabaseServer
	Db database.Database
}
