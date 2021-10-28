package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fvalladaresj/SD-Tarea2/Jugador/apiJugador"
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
				result = append(result, rand.Int31n(int32(10))+1)
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
				result = append(result, rand.Int31n(int32(10))+1)
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
			result = append(result, rand.Int31n(int32(4))+1)
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
			result = append(result, rand.Int31n(int32(10))+1)
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

type server struct {
	apiJugador.UnimplementedDataNodeJugadorServer
}

func main() {
	go manageInput()
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

func manageInput() {
	var etapa_jugada1 bool = false
	var etapa_jugada2 bool = false
	var etapa_jugada3 bool = false

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
		if response.Etapa == 1 && etapa_jugada1 == false {
			fmt.Println("Jugando primera etapa \"Luz Roja, Luz Verde\"")
			jugada := doPlay(1, false)
			response, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(1), Plays: jugada})
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
			if response.Estado[0] == 0 {
				fmt.Println("oh no! has muerto")
				etapa_jugada1 = true
				break
			}
			if checkWinner(response.Estado) {
				fmt.Println("Felicitaciones has ganado el juego del calamar")
				etapa_jugada1 = true

				break
			}
			if response.JugadorGano == 1 {
				fmt.Println("Felicitaciones por ganar la primera etapa, ingrese continuar para pasar a la siguiente etapa")
				var input string
				fmt.Scanln(&input)
				if response.Ronda < 4 {
					jugada := doPlay(1, true)
					_, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(1), Plays: jugada})
					if err != nil {
						log.Fatalf("Error Call RPC: %v", err)
					}
				}
				etapa_jugada1 = true
				fmt.Println("Espere a las instrucciones para comenzar la segunda etapa")
				continue
			}
		}

		if response.Etapa == 2 && etapa_jugada2 == false {
			fmt.Println("Jugando segunda etapa \"Tirar la cuerda\"")
			jugada := doPlay(2, false)
			response, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(2), Plays: jugada})
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
			if response.Estado[0] == 0 {
				fmt.Println("oh no! has muerto")
				break
			}
			if response.JugadorGano == 1 {
				fmt.Println("Felicitaciones por ganar la segunda etapa, ingrese continuar para pasar a la siguiente etapa")
				var input string
				fmt.Scanln(&input)
				etapa_jugada2 = true
				fmt.Println("Espere a las instrucciones para comenzar la tercera etapa")
				continue
			}
			if checkWinner(response.Estado) {
				fmt.Println("Felicitaciones has ganado el juego del calamar")
				break
			}
		}

		if response.Etapa == 3 && etapa_jugada3 == false {
			fmt.Println("Jugando tercera etapa \"Todo o Nada\"")
			jugada := doPlay(3, false)
			response, err := c.Jugar(context.Background(), &api.Jugadas{Etapa: int32(3), Plays: jugada})
			if err != nil {
				log.Fatalf("Error Call RPC: %v", err)
			}
			if response.Estado[0] == 0 {
				fmt.Println("oh no! has muerto")
				break
			}
			if response.JugadorGano == 1 {
				fmt.Println("Felicitaciones por ganar la tercera etapa, ingrese continuar para pasar a la siguiente etapa")
				var input string
				fmt.Scanln(&input)
				etapa_jugada3 = true
				fmt.Println("Felicitaciones has ganado el juego del calamar")
				break
			}
		}
	}
}

func (*server) EscribirJugada(ctx context.Context, in *apiJugador.JugadaJugador) (*apiJugador.Signal, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_Etapa string = strconv.FormatInt(int64(in.Etapa), 10)

	var str_Jugada string
	for _, jugada := range in.Jugada {
		str_Jugada = str_Jugada + strconv.FormatInt(int64(jugada), 10) + "\n"
	}

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

func (*server) RetornarJugadas(ctx context.Context, in *apiJugador.JugadorYEtapa) (*apiJugador.JugadasArchivo, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_NroEtapa string = strconv.FormatInt(int64(in.NroEtapa), 10)

	var nombre_archivo string = "jugador_" + str_Idjugador + "__ronda" + str_NroEtapa + ".txt"

	content, err := os.ReadFile(nombre_archivo)
	if err != nil {
		log.Fatal(err)
	}

	var string_content string = string(content)

	return &apiJugador.JugadasArchivo{JugadasJugador: string_content}, nil

}
