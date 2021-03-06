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

type server struct {
	apiJugador.UnimplementedDataNodeJugadorServer
}

// Funcion para que el DataNode escucha a traves de GRPC
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

// Rutina que se encarga de manejar el input del jugador
func manageInput() {
	var etapa_jugada1 bool = false
	var etapa_jugada2 bool = false
	var etapa_jugada3 bool = false

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.122:50051", grpc.WithInsecure())
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
		if response.Etapa == 1 && !etapa_jugada1 {
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
			if response.JugadorGano == 1 {
				fmt.Println("Felicitaciones por ganar la primera etapa")
				if checkWinner(response.Estado) {
					fmt.Println("Felicitaciones has ganado el juego del calamar")
					etapa_jugada1 = true
					break
				} else {
					interfaz()
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
		}

		if response.Etapa == 2 && !etapa_jugada2 {
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
			if response.Estado[0] == 1 {
				fmt.Println("Felicitaciones por ganar la segunda etapa")
				if checkWinner(response.Estado) {
					fmt.Println("Felicitaciones has ganado el juego del calamar")
					break
				} else {
					interfaz()
					etapa_jugada2 = true
					fmt.Println("Espere a las instrucciones para comenzar la tercera etapa")
					continue
				}
			}
		}

		if response.Etapa == 3 && !etapa_jugada3 {
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
			if response.Estado[0] == 1 {
				fmt.Println("Felicitaciones por ganar la tercera etapa")
				etapa_jugada3 = true
				fmt.Println("Felicitaciones has ganado el juego del calamar")
				break
			}
		}
	}
}

//Funcion que implementa la interfaz del jugador
func interfaz() {

	var dec string

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.122:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewLiderClient(conn)

	for {
		fmt.Println("Elija entre las siguientes opciones su siguiente accion a realizar")
		fmt.Println("1. Continuar a la siguiente etapa")
		fmt.Println("2. Revisar el Monto acumulado en el pozo")
		fmt.Scanln(&dec)
		if dec == "1" {
			break
		} else if dec == "2" {
			fmt.Println("El monto acumulado es:")
			response, err := c.Monto(context.Background(), &api.Signal{Sign: true})
			if err != nil {
				log.Fatalf("No se llama: %s", err)
			}
			fmt.Println(response.Monto)
		}
	}
}

/////////////////////////////////////////////// Metodos GRCP /////////////////////////////////////////////////////////////

// Funcion de DataNode: Escriba las jugadas de determinado en un archivo.
func (*server) EscribirJugada(ctx context.Context, in *apiJugador.JugadaJugador) (*apiJugador.Signal, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_Etapa string = strconv.FormatInt(int64(in.Etapa), 10)

	var str_Jugada string
	for _, jugada := range in.Jugada {
		str_Jugada = str_Jugada + strconv.FormatInt(int64(jugada), 10) + "\n"
	}

	str := []string{"jugador_", str_Idjugador, "__etapa_", str_Etapa, ".txt"}

	var nombre_archivo string = strings.Join(str, "")

	f, err := os.OpenFile(nombre_archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(str_Jugada)

	if err2 != nil {
		log.Fatal(err2)
	}

	return &apiJugador.Signal{Sign: true}, nil

}

// Esta funcion es de DataNode, retorna las jugadas de un jugador determinado de una determinada etapa segun se requiera.
func (*server) RetornarJugadas(ctx context.Context, in *apiJugador.JugadorYEtapa) (*apiJugador.JugadasArchivo, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_NroEtapa string = strconv.FormatInt(int64(in.NroEtapa), 10)

	var nombre_archivo string = "jugador_" + str_Idjugador + "__etapa_" + str_NroEtapa + ".txt"

	content, err := os.ReadFile(nombre_archivo)
	if err != nil {
		log.Fatal(err)
	}

	var string_content string = string(content)

	return &apiJugador.JugadasArchivo{JugadasJugador: string_content}, nil

}

////////////////////////////////////////////////////////// Funciones varias//////////////////////////////////////////////

// Funcion que almacena la jugada del jugador humano y genera las jugadas de los bots.
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

// funcion que Chequea si el jugador gano
func checkWinner(status []int32) bool {
	for i := 1; i < 16; i++ {
		if status[i] == 1 {
			return false
		}
	}
	return true
}
