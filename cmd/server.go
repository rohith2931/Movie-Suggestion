package main

import (
	pb "example/movieSuggestion/msproto"
	"example/movieSuggestion/schema"
	"example/movieSuggestion/server"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50501"
)

func main() {
	schema.StartDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}

	//db connection
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//create new server
	new_server := grpc.NewServer()
	pb.RegisterMsDatabaseServer(new_server, &server.MsServer{
		Db: db,
	})

	log.Printf("Using port no %v", listen.Addr())

	if err := new_server.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}
