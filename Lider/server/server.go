package server

import (
	"fmt"
	"log"
	"net"

	"github.com/fvalladaresj/SD-Tarea2/tree/main/Lider/api"
	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := api.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Lider service to the server
	api.RegisterLiderServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}




}