package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"google.golang.org/grpc"
)

var jugadores int32 = 15
var etapa_actual int32 = 0
var ronda_atual int32 = 0
var lider_e1_r1 int32 = -1
var lider_e1_r2 int32 = -1
var lider_e1_r3 int32 = -1
var lider_e1_r4 int32 = -1
var lider_e2 int32 = -1
var lider_e3 int32 = -1

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


func interfaz (decision string){
	var dec string = decision
	for {
		if (dec == "1" || dec == "2" ){
			break
		}else{
			fmt.Println("No ha ingresado una opcion valida por favor ingrese 1 o 2")
			fmt.Scanln(&dec)
		}
	}

	if (decision == "1"){
		etapa_actual = etapa_actual + 1
	}else if (decision == "2"){
		
	}else{
		
	}
}


func manageInput() {

	var input string

	fmt.Println("Bienvenido Lider, por favor espere a que hayan 16 jugadores para iniciar la partida")
	
	for {
		if !(jugadores != 16){
			break
		}
	}

	fmt.Println("Ya hay 16 Jugadores, ahora puede dar inicio a la primera etapa")
	fmt.Println("Indique el numero de la una de las siguientes acciones a realizar:")	
	fmt.Println("1. Iniciar Primera etapa")	
	fmt.Println("2. Consultar Jugadas de un jugador")			
	
	fmt.Scanln(&input)

	interfaz(input)

	
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
	}else if etapa_actual == 1 {
		return &api.State{Etapa: 1}, nil
	}else if etapa_actual == 2{
		return &api.State{Etapa: 2}, nil
	}else if etapa_actual == 3{
		return &api.State{Etapa: 3}, nil
	}else if etapa_actual == 4{
		return &api.State{Etapa: 4}, nil
	}else {
		return &api.State{Etapa: 9}, nil
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