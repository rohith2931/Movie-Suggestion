package server

import (
	"context"
	"example/movieSuggestion/pkg/database"
	pb "example/movieSuggestion/msproto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func MockServer(mockDb *database.MockDatabase, ctx context.Context) (pb.MsDatabaseClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	s := MsServer{
		Db: mockDb,
	}
	pb.RegisterMsDatabaseServer(baseServer, &s)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewMsDatabaseClient(conn)

	return client, closer
}
