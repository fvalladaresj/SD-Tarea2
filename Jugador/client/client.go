package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"google.golang.org/grpc"
)

func doPlay(etapa int) [16]int {
	var result [16]int
	var jugada int32

	if etapa == 1 {
		fmt.Println("Por favor ingrese un numero del 1 al 10")
		fmt.Scanln(&jugada)
		for{
			if (jugada >= 1 && jugada <= 10){
				break
			}else {
				fmt.Println("Por favor ingrese un numero del 1 al 10")
				fmt.Scanln(&jugada)
			}
		}
		result[0] = jugada
		for i := 1; i < 16; i++ {
			result[i] = rand.Intn(9) + 1
		}
	}
	return result
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewLiderClient(conn)

	var jugada [16]int
	var input string

	fmt.Println("Para participar en el Juego del calamar ingrese participar ")
	for {
		fmt.Scanln(&input)
		response, err := c.ParticiparJuego(context.Background(), &api.PeticionParticipar{Participar: input})
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}
		if !(response.Confirmacion != "ya esta participando en el juego del calamar") {
			fmt.Println(response)
			break

		} else {
			fmt.Println(response)
		}
	}

	fmt.Println("Espere a las instrucciones para comenzar primera etapa")

	for {
		response, err := c.EstadoEtapas(context.Background(), &api.Check{Sign: true})
		if err != nil {
			log.Fatalf("Error Call RPC: %v", err)
		}
		if !(response.Etapa = 1) {
			break
		}
	}
	fmt.Println("Ha comenzado la primera etapa, Luz Verde, Luz Roja, ingrese un numero entre el 1 y el 10")
	
	jugada = doPlay(1)
	
	response, err = c.Jugar(context.Background(), &api.Jugadas{jugadas: jugada })

}
