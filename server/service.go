package server

import (
	"example/movieSuggestion/database"
	pb "example/movieSuggestion/msproto"
)

// This struct implements all the RPC's
type MsServer struct {
	pb.UnimplementedMsDatabaseServer
	Db database.Database
}
