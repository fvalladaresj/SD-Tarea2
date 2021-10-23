package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"google.golang.org/grpc"
)

type server struct {
	api.UnimplementedLiderServer
}

// main start a gRPC server and waits for connection
func main() {
	// listen to input concurrently
	go manageInput()
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := grpc.NewServer()
	// attach the Lider service to the server
	api.RegisterLiderServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

var jugadores int32 = 15
var etapa_actual int32 = 0
var ronda_atual int32 = 0
var lider_e1_r1 int32 = -1
var lider_e1_r2 int32 = -1
var lider_e1_r3 int32 = -1
var lider_e1_r4 int32 = -1
var lider_e2 int32 = -1
var lider_e3 int32 = -1

func manageInput() {
	for {
		input := ""
		fmt.Println("ingrese algo:")
		fmt.Scanln(&input)
		fmt.Println(input)
	}
}

func (*server) ParticiparJuego(ctx context.Context, in *api.PeticionParticipar) (*api.ConfirmacionParticipacion, error) {
	if in.Participar == "participar" {
		if jugadores < 16 {

			jugadores = jugadores + 1
			return &api.ConfirmacionParticipacion{Confirmacion: "ya esta participando en el juego del calamar"}, nil
		} else {

			return &api.ConfirmacionParticipacion{Confirmacion: "ya no puede participar, el juego esta lleno"}, nil
		}
	} else {

		return &api.ConfirmacionParticipacion{Confirmacion: "verifique que ha escrito bien participar"}, nil
	}
}

func (*server) EstadoEtapas(ctx context.Context, in *api.Check) (*api.State, error) {
	if etapa_actual == 0 {
		return &api.State{Etapa: 0}, nil
	}
	if etapa_actual == 1 {
		return &api.State{Etapa: 1}, nil
	} else {
		return &api.State{Etapa: 1}, nil
	}
}

func (*server) CuantosJugadores(ctx context.Context, in *api.Signal) (*api.CantidadJugadores, error) {
	return &api.CantidadJugadores{CJugadores: jugadores}, nil
}

func (*server) Iniciar(ctx context.Context, in *api.Signal) (*api.Signal, error) {
	etapa_actual = etapa_actual + 1
	return &api.Signal{Sign: true}, nil
}

/*
func (*server) Jugar(ctx context.Context, in *api.Jugadas) (*api.EstadoJugador, error) {
}

func (*server) Monto(ctx context.Context, in *api.PedirMonto) (*api.MontoJugador, error) {

}
*/
