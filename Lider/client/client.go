package main

import(
	"context"
	"fmt"
	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"golang.org/x/net/context"
  	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
	  log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewLiderClient(conn)

	var input string

	fmt.Println("Bienvenido Lider, por favor espere a que hayan 16 jugadores para iniciar la partida")
	
	for {
		response, err = c.CuantosJugadores(context.Background(), &api.Check{Sign = 1})
		if !(response != 16){
			break
		}
		else{
			fmt.Println("Ya hay 16 Jugadores, ahora puede dar inicio a la primera etapa")
		}
	}

	fmt.Scanln(&input)
	


}