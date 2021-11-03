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

	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"github.com/fvalladaresj/SD-Tarea2/NameNode/apiNameNode"
	"github.com/fvalladaresj/SD-Tarea2/Pozo/apiPozo"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

// almacena la cuenta de la cantidad de jugadores que hay en el juego, comienza en 15 porque son bots.
var jugadores int32 = 15

// almacena el estado de los jugadores del juego
var est_jugadores []int32 = []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

// almacena la etapa actual en la que va el juego del calamar
var etapa_actual int32 = 0

// sirve para controlar el flujo de las etapas y cuando se tienen que jugar
var etapa_check_1 bool = true
var etapa_check_2 bool = false
var etapa_check_3 bool = false
var etapa_check_4 bool = false

// almacena ronda actual de la etapa 1 del juego del calamar
var rnd_actual int32 = 0

// almacena la puntuacion en la etapa 1 de cada jugador
var pts_jugadores_e1 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// almacena a los ganadores de la etapa 1
var ganadores_e1 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

// Almacena todas las jugadas que realizan los jugadores durante la etapa 1
var jugadas_acumuladas_e1 [][]int32

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

// rutina que permite tener la interfaz del lider
func manageInput() {
	var input string
	fmt.Println("Bienvenido Lider, por favor espere a que hayan 16 jugadores para iniciar la partida")
	for {
		if !(jugadores != 16) {
			break
		}
	}

	for {
		if est_jugadores[0] == 0 {
			if etapa_actual > 1 {
				etapa_check_2 = true
			}
			if checkPlayers() == "Ganador" || etapa_actual > 3 {
				fmt.Println("El juego del calamar ha finalizado, felicitaciones a los ganadores:")
				PrintAlive()
				break
			} else {
				fmt.Println("El jugador principal ha muerto, procediendo automáticamente")
				for {
					if checkPlayers() == "Ganador" || etapa_actual > 3 || etapa_check_4 {
						fmt.Println("El juego del calamar ha finalizado, felicitaciones a los ganadores:")
						PrintAlive()
						break
					} else if checkPlayers() == "Muertos" {
						fmt.Println("Todos los jugadores han muerto")
						break
					} else {
						if etapa_actual == 1 && etapa_check_2 {
							etapa_check_2 = false
							for {
								fmt.Println("Indique el numero de una de las siguientes accione2s a realizar:")
								fmt.Println("1. Iniciar la segunda etapa")
								fmt.Println("2. Consultar Jugadas de un jugador")
								fmt.Scanln(&input)
								if input == "1" {
									interfaz(input)
									moves := doPlay(int(etapa_actual))
									JugarAutomatico(moves, etapa_actual)
									break
								} else {
									interfaz(input)
								}
							}
						} else if etapa_actual == 2 && etapa_check_3 {
							etapa_check_3 = false
							for {
								fmt.Println("Indique el numero de una de las siguientes acciones a realizar:")
								fmt.Println("1. Iniciar la tercera etapa")
								fmt.Println("2. Consultar Jugadas de un jugador")
								fmt.Scanln(&input)
								if input == "1" {
									interfaz(input)
									moves := doPlay(int(etapa_actual))
									JugarAutomatico(moves, etapa_actual)
									break
								} else {
									interfaz(input)
								}
							}
						} else {
							moves := doPlay(int(etapa_actual))
							JugarAutomatico(moves, etapa_actual)
						}
					}
				}
			}
			break
		} else {
			if etapa_actual == 0 && etapa_check_1 {
				etapa_check_1 = false
				fmt.Println("Ya hay 16 Jugadores, ahora puede dar inicio a la primera etapa")
				for {
					fmt.Println("Indique el numero de una de las siguientes acciones a realizar:")
					fmt.Println("1. Iniciar la primera etapa")
					fmt.Println("2. Consultar Jugadas de un jugador")
					fmt.Scanln(&input)
					if input == "1" {
						interfaz(input)
						break
					} else {
						interfaz(input)
					}
				}
			} else if etapa_actual == 1 && etapa_check_2 {
				etapa_check_2 = false
				for {
					fmt.Println("Indique el numero de una de las siguientes acciones a realizar:")
					fmt.Println("1. Iniciar la segunda etapa")
					fmt.Println("2. Consultar Jugadas de un jugador")
					fmt.Scanln(&input)
					if input == "1" {
						interfaz(input)
						break
					} else {
						interfaz(input)
					}
				}
			} else if etapa_actual == 2 && etapa_check_3 {
				etapa_check_3 = false
				for {
					fmt.Println("Indique el numero de una de las siguientes acciones a realizar:")
					fmt.Println("1. Iniciar la tercera etapa")
					fmt.Println("2. Consultar Jugadas de un jugador")
					fmt.Scanln(&input)
					if input == "1" {
						interfaz(input)
						break
					} else {
						interfaz(input)
					}
				}
			} else if etapa_actual == 3 && etapa_check_4 {
				break
			}
		}
	}
}

// Administra las deciciones que puede hacer el lider en su interfaz
func interfaz(decision string) {
	var dec string = decision

	for {
		if dec == "1" || dec == "2" {
			break
		} else {
			fmt.Println("No ha ingresado una opcion valida por favor ingrese 1 o 2")
			break
			// fmt.Scanln(&dec)
		}
	}

	if decision == "1" {
		etapa_actual = etapa_actual + 1
	} else if decision == "2" {
		if etapa_actual < 1 {
			fmt.Println("Aún no se ha jugado la primera etapa por lo que no es posible revisar jugadas.")
		} else {
			id := 0
			fmt.Println("Ingrese id del jugador: ")
			fmt.Scanln(&id)
			fmt.Println(BuscarJugadas(int32(id)))
		}
	}

}

/////////////////////////////////////////////// Metodos GRCP /////////////////////////////////////////////////////////////

// Esta funcion sirve para inscribirse en el juego del calamar.
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

// Informa la etapa actual en la que esta el Juego
func (*server) EstadoEtapas(ctx context.Context, in *api.Check) (*api.State, error) {
	if etapa_actual == 0 {
		return &api.State{Etapa: 0}, nil
	} else if etapa_actual == 1 {
		return &api.State{Etapa: 1}, nil
	} else if etapa_actual == 2 {
		return &api.State{Etapa: 2}, nil
	} else if etapa_actual == 3 {
		return &api.State{Etapa: 3}, nil
	} else if etapa_actual == 4 {
		return &api.State{Etapa: 4}, nil
	} else {
		return &api.State{Etapa: 9}, nil
	}
}

// Retorna la cantidad de jugadores que hay en el juego.
func (*server) CuantosJugadores(ctx context.Context, in *api.Signal) (*api.CantidadJugadores, error) {
	return &api.CantidadJugadores{CJugadores: jugadores}, nil
}

// Esta funcion Recibe las jugadas de los jugadores y lleva gran parte de la logica del juego, definiendo lo que se realiza segun la etapa.
func (*server) Jugar(ctx context.Context, in *api.Jugadas) (*api.EstadoJugador, error) {
	moves := in.Plays
	if in.Etapa == 1 {
		players := canPlayPhase1()
		for _, player := range players {
			if len(jugadas_acumuladas_e1) < 16 {
				aux := []int32{moves[player]}
				jugadas_acumuladas_e1 = append(jugadas_acumuladas_e1, aux)
			} else {
				jugadas_acumuladas_e1[player] = append(jugadas_acumuladas_e1[player], moves[player])
			}
		}

		if rnd_actual < 3 && len(players) > 0 {
			rand.Seed(time.Now().UnixNano())

			leaderMove := rand.Int31n(int32(4)) + int32(6)
			for _, player := range players {
				if moves[player] >= leaderMove {
					est_jugadores[player] = 0 //muerto
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				} else {
					pts_jugadores_e1[player] = pts_jugadores_e1[player] + moves[player]
					if pts_jugadores_e1[player] >= 21 {
						ganadores_e1[player] = 1
					}
				}
			}
			rnd_actual = rnd_actual + 1

			return &api.EstadoJugador{Estado: est_jugadores, Ronda: rnd_actual, JugadorGano: ganadores_e1[0]}, nil
		} else {
			for _, player := range players {
				pts_jugadores_e1[player] = pts_jugadores_e1[player] + moves[player]
				if pts_jugadores_e1[player] < 21 {
					est_jugadores[player] = 0
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				} else {
					ganadores_e1[player] = 1
				}
			}
			fmt.Println("Etapa 1 finalizada, jugadores vivos: ")
			PrintAlive()
			etapa_check_2 = true
			EscribirNameNodeEtapa1()
			return &api.EstadoJugador{Estado: est_jugadores, Ronda: 4, JugadorGano: ganadores_e1[0]}, nil
		}
	} else if in.Etapa == 2 {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(4)) + int32(1)
		players := canPlayPhase2()
		if len(players)%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players)))
			est_jugadores[players[indexToDelete]] = 0 //muerto
			fmt.Printf("Jugador %v ha muerto \n", players[indexToDelete])
			go sendRabbit(players[indexToDelete], etapa_actual)
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		teamA := players[0 : len(players)/2]
		teamB := players[len(players)/2:]

		if sum(teamA)%2 == sum(teamB)%2 && sum(teamA)%2 != leaderMove%2 {
			if rand.Int31n(int32(2)) == 0 {
				for _, player := range teamA {
					est_jugadores[player] = 0
				}
			} else {
				for _, player := range teamB {
					est_jugadores[player] = 0
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				}
			}
		} else if sum(teamA)%2 != leaderMove%2 {
			for _, player := range teamA {
				est_jugadores[player] = 0
				fmt.Printf("Jugador %v ha muerto \n", player)
				go sendRabbit(player, etapa_actual)
			}
		} else if sum(teamB)%2 != leaderMove%2 {
			for _, player := range teamB {
				est_jugadores[player] = 0
				fmt.Printf("Jugador %v ha muerto \n", player)
				go sendRabbit(player, etapa_actual)
			}
		}
		fmt.Println("Etapa 2 finalizada, jugadores vivos: ")
		PrintAlive()
		etapa_check_3 = true
		EscribirNameNodeEtapa2y3(int32(2), players, moves)
		return &api.EstadoJugador{Estado: est_jugadores}, nil
	} else {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(4)) + int32(1)
		players := canPlayPhase2()
		if len(canPlayPhase2())%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players)))
			est_jugadores[players[indexToDelete]] = 0 //muerto
			fmt.Printf("Jugador %v ha muerto \n", players[indexToDelete])
			go sendRabbit(indexToDelete, etapa_actual)
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		playerCouples := tuples(players)
		for _, couple := range playerCouples {
			if couple[0] != couple[1] {
				if Abs(couple[0]-leaderMove) > Abs(couple[1]-leaderMove) {
					est_jugadores[couple[0]] = 0
					fmt.Printf("Jugador %v ha muerto \n", couple[0])
					go sendRabbit(couple[0], etapa_actual)

				} else {
					est_jugadores[couple[1]] = 0
					fmt.Printf("Jugador %v ha muerto \n", couple[1])
					go sendRabbit(couple[1], etapa_actual)

				}
			}
		}
		fmt.Println("El juego del calamar ha finalizado, felicitaciones a los ganadores:")
		PrintAlive()
		etapa_check_4 = true
		EscribirNameNodeEtapa2y3(int32(3), players, moves)
		return &api.EstadoJugador{Estado: est_jugadores}, nil
	}
}

// Funcion de DataNode: Escriba las jugadas de determinado en un archivo.
func (*server) EscribirJugada(ctx context.Context, in *api.JugadaJugador) (*api.Signal, error) {
	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_Etapa string = strconv.FormatInt(int64(in.Etapa), 10)

	var str_Jugada string
	for _, jugada := range in.Jugada {
		str_Jugada = str_Jugada + strconv.FormatInt(int64(jugada), 10) + "\n"
	}

	str := []string{"jugador_", str_Idjugador, "__ronda_", str_Etapa, ".txt"}

	var nombre_archivo string = strings.Join(str, "")

	f, err := os.OpenFile(nombre_archivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(str_Jugada)

	if err2 != nil {
		log.Fatal(err2)
	}

	return &api.Signal{Sign: true}, nil

}

// Esta funcion es de DataNode, retorna las jugadas de un jugador determinado de una determinada etapa segun se requiera.
func (*server) RetornarJugadas(ctx context.Context, in *api.JugadorYEtapa) (*api.JugadasArchivo, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_NroEtapa string = strconv.FormatInt(int64(in.NroEtapa), 10)

	var nombre_archivo string = "jugador_" + str_Idjugador + "__ronda_" + str_NroEtapa + ".txt"

	content, err := os.ReadFile(nombre_archivo)
	if err != nil {
		log.Fatal(err)
	}

	var string_content string = string(content)

	return &api.JugadasArchivo{JugadasJugador: string_content}, nil

}

// esta funcion sirve para retornar el monto acumulado en el pozo
func (*server) Monto(ctx context.Context, in *api.Signal) (*api.MontoJugador, error) {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.124:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := apiPozo.NewDataNodePozoClient(conn)

	response, err := c.Monto(context.Background(), &apiPozo.Signal{Sign: true})
	if err != nil {
		log.Fatalf("Error Call RPC: %v", err)
	}

	return &api.MontoJugador{Monto: response.Monto}, nil

}

////////////////////////////////////////////////////////// Funciones varias//////////////////////////////////////////////

// chequea la cantidad de jugadores vivos que hay.
func checkPlayers() string {
	counter := 0
	for _, playerStatus := range est_jugadores {
		if playerStatus == 1 {
			counter += 1
		}
	}
	if counter == 0 {
		return "Muertos"
	} else if counter > 1 {
		return "Jugable"
	} else {
		return "Ganador"
	}
}

// Esta funcion lleva la logica de los juegos automaticos de los bots.
func JugarAutomatico(moves []int32, etapa int32) {
	if etapa == 1 {
		players := canPlayPhase1()
		for _, player := range players {
			if len(jugadas_acumuladas_e1) < 16 {
				aux := []int32{moves[player]}
				jugadas_acumuladas_e1 = append(jugadas_acumuladas_e1, aux)
			} else {
				jugadas_acumuladas_e1[player] = append(jugadas_acumuladas_e1[player], moves[player])
			}
		}
		if rnd_actual < 3 && len(players) > 0 {
			rand.Seed(time.Now().UnixNano())

			leaderMove := rand.Int31n(int32(5)) + int32(6)
			for _, player := range players {
				if moves[player] >= leaderMove {
					est_jugadores[player] = 0 //muerto
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				} else {
					pts_jugadores_e1[player] = pts_jugadores_e1[player] + moves[player]
					if pts_jugadores_e1[player] >= 21 {
						ganadores_e1[player] = 1
					}
				}
			}
			rnd_actual = rnd_actual + 1
		} else {
			for _, player := range players {
				pts_jugadores_e1[player] = pts_jugadores_e1[player] + moves[player]
				if pts_jugadores_e1[player] < 21 {
					est_jugadores[player] = 0
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				} else {
					ganadores_e1[player] = 1
				}
			}
			fmt.Print("Etapa 1 finalizada, jugadores vivos: ")
			PrintAlive()
			etapa_check_2 = true
			EscribirNameNodeEtapa1()
		}
	} else if etapa == 2 {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(4)) + int32(1)
		players := canPlayPhase2()
		if len(players)%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players)))
			est_jugadores[players[indexToDelete]] = 0 //muerto
			fmt.Printf("Jugador %v ha muerto \n", players[indexToDelete])
			go sendRabbit(players[indexToDelete], etapa_actual)
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		teamA := players[0 : len(players)/2]
		teamB := players[len(players)/2:]

		if sum(teamA)%2 == sum(teamB)%2 && sum(teamA)%2 != leaderMove%2 {
			if rand.Int31n(int32(2)) == 0 {
				for _, player := range teamA {
					est_jugadores[player] = 0
				}
			} else {
				for _, player := range teamB {
					est_jugadores[player] = 0
					fmt.Printf("Jugador %v ha muerto \n", player)
					go sendRabbit(player, etapa_actual)
				}
			}
		} else if sum(teamA)%2 != leaderMove%2 {
			for _, player := range teamA {
				est_jugadores[player] = 0
				fmt.Printf("Jugador %v ha muerto \n", player)
				go sendRabbit(player, etapa_actual)
			}
		} else if sum(teamB)%2 != leaderMove%2 {
			for _, player := range teamB {
				est_jugadores[player] = 0
				fmt.Printf("Jugador %v ha muerto \n", player)
				go sendRabbit(player, etapa_actual)
			}
		}
		fmt.Print("Etapa 2 finalizada, jugadores vivos: ")
		PrintAlive()
		etapa_check_3 = true
		EscribirNameNodeEtapa2y3(int32(2), players, moves)
	} else {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(4)) + int32(1)
		players := canPlayPhase2()
		if len(canPlayPhase2())%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players)))
			est_jugadores[players[indexToDelete]] = 0 //muerto
			fmt.Printf("Jugador %v ha muerto \n", players[indexToDelete])
			go sendRabbit(indexToDelete, etapa_actual)
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		playerCouples := tuples(players)
		for _, couple := range playerCouples {
			if couple[0] != couple[1] {
				if Abs(couple[0]-leaderMove) > Abs(couple[1]-leaderMove) {
					est_jugadores[couple[0]] = 0
					fmt.Printf("Jugador %v ha muerto \n", couple[0])
					go sendRabbit(couple[0], etapa_actual)

				} else {
					est_jugadores[couple[1]] = 0
					fmt.Printf("Jugador %v ha muerto \n", couple[1])
					go sendRabbit(couple[1], etapa_actual)

				}
			}
		}
		fmt.Println("Etapa 3 finalizada, jugadores vivos: ")
		PrintAlive()
		etapa_check_4 = true
		EscribirNameNodeEtapa2y3(int32(3), players, moves)
	}
}

// Determina si los jugadores pueden seguir jugando las rondas de la etapa 1

func canPlayPhase1() []int32 {
	var result []int32
	for i := range est_jugadores {
		// si esta vivo y no ha ganado la etapa
		if est_jugadores[i] == 1 && ganadores_e1[i] == 0 {
			result = append(result, int32(i))
		}
	}
	return result
}

// Determina si un jugador puede jugar la etapa 2 y 3

func canPlayPhase2() []int32 {
	var result []int32
	for i := range est_jugadores {
		// si esta vivo y no ha ganado la etapa
		if est_jugadores[i] == 1 {
			result = append(result, int32(i))
		}
	}
	return result
}

func sum(array []int32) int32 {
	result := int32(0)
	for _, v := range array {
		result += v
	}
	return result
}

func tuples(array []int32) [][]int32 {
	var result [][]int32
	index := 0
	for index < len(array) {
		result = append(result, array[index:index+2])
		index = index + 2
	}
	return result
}

// Retorna valor absoluto de un numero dado
func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Manejo de Errores
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Funcion que se encarga de reportar cuando mueren jugadores al Poz
func sendRabbit(player int32, round int32) {
	conn, err := amqp.Dial("amqp:/usuario:guamelo1324@10.6.43.120:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Jugadores muertos", // name
		false,               // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Jugador_" + strconv.FormatInt(int64(player), 10) + " " + "Ronda_" + strconv.FormatInt(int64(round), 10)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

// Escribe las jugadas realizadas por los jugadores en la etapa 1

func EscribirNameNodeEtapa1() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.123:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := apiNameNode.NewNameNodeClient(conn)
	for player, moves := range jugadas_acumuladas_e1 {
		c.EscribirJugada(context.Background(), &apiNameNode.JugadaJugador{IdJugador: int32(player), Jugada: moves, Etapa: int32(1)})
	}
}

// Escribe las jugadas realizadas por los jugadores en la etapa 2 y 3
func EscribirNameNodeEtapa2y3(etapa int32, jugadores []int32, jugadas []int32) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.123:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := apiNameNode.NewNameNodeClient(conn)
	for _, player := range jugadores {
		aux := []int32{jugadas[player]}
		c.EscribirJugada(context.Background(), &apiNameNode.JugadaJugador{IdJugador: int32(player), Jugada: aux, Etapa: etapa})
	}
}

// Se encarga de ir a buscar las jugadas de determinado jugador al NameNode.
func BuscarJugadas(id int32) string {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.6.43.123:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := apiNameNode.NewNameNodeClient(conn)
	response, err := c.PedirJugadasJugador(context.Background(), &apiNameNode.Jugador{IdJugador: int32(id)})
	if err != nil {
		log.Fatal(err)
	}
	return response.JugadasJugador
}

// Genera un array con las jugadas de todos los jugadores
func doPlay(etapa int) []int32 {
	var result []int32

	rand.Seed(time.Now().UnixNano())

	if etapa == 1 {
		for i := 0; i < 16; i++ {
			result = append(result, rand.Int31n(int32(10))+1)
		}
	} else if etapa == 2 {
		for i := 0; i < 16; i++ {
			result = append(result, rand.Int31n(int32(4))+1)
		}
	} else if etapa == 3 {
		for i := 0; i < 16; i++ {
			result = append(result, rand.Int31n(int32(10))+1)
		}
	}
	return result
}

// Imprime los jugadores que estan vivos cuando se invoca
func PrintAlive() {
	for i := range est_jugadores {
		if est_jugadores[i] == 1 {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}
