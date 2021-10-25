package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/fvalladaresj/SD-Tarea2/Jugador/apiJugador"
	"google.golang.org/grpc"
)

type server struct {
	apiJugador.UnimplementedDataNodeJugadorServer
}

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 50053
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := grpc.NewServer()
	// attach the Lider service to the server
	apiJugador.RegisterDataNodeJugadorServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func (*server) EscribirJugada(ctx context.Context, in *apiJugador.JugadaJugador) (*apiJugador.Signal, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_Jugada string = strconv.FormatInt(int64(in.Jugada), 10)
	var str_Etapa string = strconv.FormatInt(int64(in.Etapa), 10)

	str := []string{"jugador_", str_Idjugador, "__ronda", str_Etapa, ".txt"}

	var nombre_archivo string = strings.Join(str, "")

	f, err := os.OpenFile(nombre_archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(str_Jugada + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	return &apiJugador.Signal{Sign: true}, nil

}
