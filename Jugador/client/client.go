package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"google.golang.org/grpc"
)

func doPlay(etapa int, gano bool) []int32 {
	var result []int32
	var jugada int32

	rand.Seed(time.Now().UnixNano())

	if etapa == 1 {
		if gano {
			result = append(result, int32(0))
			for i := 0; i < 15; i++ {
				result = append(result, rand.Int31n(int32(9))+1)
			}
		} else {
			fmt.Println("Por favor ingrese un numero del 1 al 10")
			fmt.Scanln(&jugada)
			for {
				if jugada >= 1 && jugada <= 10 {
					break
				} else {
					fmt.Println("Por favor ingrese un numero del 1 al 10")
					fmt.Scanln(&jugada)
				}
			}
			result = append(result, int32(jugada))
			for i := 0; i < 15; i++ {
				result = append(result, rand.Int31n(int32(9))+1)
			}
		}
	} else if etapa == 2 {
		fmt.Println("Por favor ingrese un numero del 1 al 4")
		fmt.Scanln(&jugada)
		for {
			if jugada >= 1 && jugada <= 4 {
				break
			} else {
				fmt.Println("Por favor ingrese un numero del 1 al 4")
				fmt.Scanln(&jugada)
			}
		}
		result = append(result, int32(jugada))
		for i := 0; i < 15; i++ {
			result = append(result, rand.Int31n(int32(3))+1)
		}
	} else if etapa == 3 {
		fmt.Println("Por favor ingrese un numero del 1 al 10")
		fmt.Scanln(&jugada)
		for {
			if jugada >= 1 && jugada <= 10 {
				break
			} else {
				fmt.Println("Por favor ingrese un numero del 1 al 10")
				fmt.Scanln(&jugada)
			}
		}
		result = append(result, int32(jugada))
		for i := 0; i < 15; i++ {
			result = append(result, rand.Int31n(int32(9))+1)
		}
	}
	return result
}

func checkWinner(status []int32) bool {
	for i := 1; i < 16; i++ {
		if status[i] == 1 {
			return false
		}
	}
	return true
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewLiderClient(conn)

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
		if response.Etapa == 1 {
			fmt.Println("Jugando primera etapa \"Luz Roja, Luz Verde\"")
			jugada := doPlay(1, false)
			response, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(1), Plays: jugada})
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
			if response.Estado[0] == 0 {
				fmt.Println("oh no! has muerto")
				break
			}
			if checkWinner(response.Estado) {
				fmt.Println("Felicitaciones has ganado el juego del calamar")
				break
			} else if response.JugadorGano == 1 {
				fmt.Println("Felicitaciones por ganar la primera etapa, ingrese continuar para pasar a la siguiente etapa")
				input := ' '
				fmt.Scanln(&input)
				if response.Ronda < 4 {
					jugada := doPlay(1, true)
					_, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(1), Plays: jugada})
					if err != nil {
						log.Fatalf("Error Call RPC: %v", err)
					}
				}
			}
		}
	}
}
