package server

import (
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/pkg/database"
)

// This struct implements all the RPC's
type MsServer struct {
	pb.UnimplementedMovieSuggestionServiceServer
	Db database.Database
}
