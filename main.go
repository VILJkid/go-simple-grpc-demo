package main

import (
	"log"
	"net"

	"github.com/VILJkid/go-simple-grpc-demo/proto"
	"github.com/VILJkid/go-simple-grpc-demo/service"
	"google.golang.org/grpc"
)

const port = ":8081"

func main() {
	// Create a new gRPC server and register the User service
	server := grpc.NewServer()
	userServiceServer := service.NewUserServiceServer()
	proto.RegisterUserServiceServer(server, userServiceServer)

	// Start the server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start listener %v", err)
	}
	log.Printf("Server started at %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
