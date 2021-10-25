package main

import (
	"log"
	"net"

	"github.com/fvalladaresj/SD-Tarea2/NameNode/apiNameNode"

	"google.golang.org/grpc"
)

type server struct {
	apiNameNode.UnimplementedNameNodeServer
}

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 50052
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := grpc.NewServer()
	// attach the Lider service to the server
	apiNameNode.RegisterNameNodeServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
