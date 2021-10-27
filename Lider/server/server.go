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
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var jugadores int32 = 15
var est_jugadores []int32 = []int32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var etapa_actual int32 = 0

var etapa_check_1 bool = true
var etapa_check_2 bool = false
var etapa_check_3 bool = false

var rnd_actual int32 = 0
var pts_jugadores_e1 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var ganadores_e1 [16]int32 = [16]int32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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

func interfaz(decision string) {

	var dec string = decision

	for {
		if dec == "1" || dec == "2" {
			break
		} else {
			fmt.Println("No ha ingresado una opcion valida por favor ingrese 1 o 2")
			fmt.Scanln(&dec)
		}
	}

	if decision == "1" {
		etapa_actual = etapa_actual + 1
	} else if decision == "2" {

	} else {

	}
}

func manageInput() {

	var input string

	fmt.Println("Bienvenido Lider, por favor espere a que hayan 16 jugadores para iniciar la partida")
	for {
		if !(jugadores != 16) {
			break
		}
	}

	for {

		if etapa_actual == 0 && etapa_check_1 {
			etapa_check_1 = false
			fmt.Println("Ya hay 16 Jugadores, ahora puede dar inicio a la primera etapa")
			fmt.Println("Indique el numero de la una de las siguientes acciones a realizar:")
			fmt.Println("1. Iniciar primera etapa")
			fmt.Println("2. Consultar Jugadas de un jugador")
			fmt.Scanln(&input)
			interfaz(input)
		} else if etapa_actual == 1 && etapa_check_2 {
			etapa_check_2 = false
			fmt.Println("Indique el numero de la una de las siguientes acciones a realizar:")
			fmt.Println("1. Iniciar la segunda etapa")
			fmt.Println("2. Consultar Jugadas de un jugador")
			fmt.Scanln(&input)
			interfaz(input)
		} else if etapa_actual == 3 && etapa_check_3 {
			etapa_check_3 = false
			fmt.Println("Indique el numero de la una de las siguientes acciones a realizar:")
			fmt.Println("1. Iniciar la segunda etapa")
			fmt.Println("2. Consultar Jugadas de un jugador")
			fmt.Scanln(&input)
			interfaz(input)
		}
	}
}

/////////////////////////////////////////////// Metodos GRCP /////////////////////////////////////////////////////////////

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

func (*server) CuantosJugadores(ctx context.Context, in *api.Signal) (*api.CantidadJugadores, error) {
	return &api.CantidadJugadores{CJugadores: jugadores}, nil
}

func (*server) Jugar(ctx context.Context, in *api.Jugadas) (*api.EstadoJugador, error) {
	moves := in.Plays
	if in.Etapa == 1 {
		players := canPlayPhase1()
		if rnd_actual < 3 && len(players) > 0 {
			rand.Seed(time.Now().UnixNano())

			leaderMove := rand.Int31n(int32(4)) + int32(6)
			//log.Printf("Lider: %v", leaderMove)
			for _, player := range players {
				if moves[player] >= leaderMove {
					est_jugadores[player] = 0 //muerto
					//log.Printf("Jugador %v ha muerto, tiro %v y tiene %v puntos", player, moves[player], pts_jugadores_e1[player])
					log.Printf("Jugador %v ha muerto", player)
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
					//log.Printf("Jugador %v ha muerto, tiro %v y tiene %v puntos", player, moves[player], pts_jugadores_e1[player])
					log.Printf("Jugador %v ha muerto", player)
					go sendRabbit(player, etapa_actual)
				} else {
					ganadores_e1[player] = 1
				}
			}
			fmt.Print("Etapa 1 finalizada, jugadores vivos: ")
			for i := range est_jugadores {
				if est_jugadores[i] == 1 {
					fmt.Printf("%v ", i)
				}
			}
			etapa_check_2 = true
			fmt.Println()
			return &api.EstadoJugador{Estado: est_jugadores, Ronda: 4, JugadorGano: ganadores_e1[0]}, nil
		}
	} else if in.Etapa == 2 {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(3)) + int32(1)
		players := canPlayPhase2()
		if len(players)%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players))) + int32(1)
			est_jugadores[indexToDelete] = 0 //muerto
			log.Printf("Jugador %v ha muerto", players[indexToDelete])
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		teamA := players[0 : len(players)/2]
		teamB := players[len(players)/2:]

		if sum(teamA)%2 == sum(teamB)%2 && sum(teamA)%2 != leaderMove%2 {
			if rand.Int31n(int32(1)) == 0 {
				for _, player := range teamA {
					est_jugadores[player] = 0
					log.Printf("Jugador %v ha muerto", player)
				}
			} else {
				for _, player := range teamB {
					est_jugadores[player] = 0
					log.Printf("Jugador %v ha muerto", player)
				}
			}
		} else if sum(teamA)%2 != leaderMove%2 {
			for _, player := range teamA {
				est_jugadores[player] = 0
				log.Printf("Jugador %v ha muerto", player)
			}
		} else if sum(teamB)%2 != leaderMove%2 {
			for _, player := range teamB {
				est_jugadores[player] = 0
				log.Printf("Jugador %v ha muerto", player)
			}
		}
		return &api.EstadoJugador{Estado: est_jugadores, JugadorGano: ganadores_e1[0]}, nil
	} else {
		rand.Seed(time.Now().UnixNano())
		leaderMove := rand.Int31n(int32(3)) + int32(1)
		players := canPlayPhase2()
		if len(canPlayPhase2())%2 == 1 { //es impar
			indexToDelete := rand.Int31n(int32(len(players))) + int32(1)
			est_jugadores[indexToDelete] = 0 //muerto
			log.Printf("Jugador %v ha muerto", players[indexToDelete])
			players = append(players[:indexToDelete], players[indexToDelete+1:]...)
		}
		playerCouples := tuples(players)
		for _, couple := range playerCouples {
			if couple[0] != couple[1] {
				if Abs(couple[0]-leaderMove) > Abs(couple[1]-leaderMove) {
					est_jugadores[couple[0]] = 0
					log.Printf("Jugador %v ha muerto", couple[0])
				} else {
					est_jugadores[couple[1]] = 0
					log.Printf("Jugador %v ha muerto", couple[1])
				}
			}
		}
		return &api.EstadoJugador{Estado: est_jugadores, JugadorGano: ganadores_e1[0]}, nil
	}
}

func (*server) EscribirJugada(ctx context.Context, in *api.JugadaJugador) (*api.Signal, error) {

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

	return &api.Signal{Sign: true}, nil

}

func (*server) RetornarJugadas(ctx context.Context, in *api.JugadorYEtapa) (*api.JugadasArchivo, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_NroEtapa string = strconv.FormatInt(int64(in.NroEtapa), 10)

	var nombre_archivo string = "jugador_" + str_Idjugador + "__ronda" + str_NroEtapa + ".txt"

	content, err := os.ReadFile(nombre_archivo)
	if err != nil {
		log.Fatal(err)
	}

	var string_content string = string(content)

	return &api.JugadasArchivo{JugadasJugador: string_content}, nil

}

/*
func (*server) Monto(ctx context.Context, in *api.PedirMonto) (*api.MontoJugador, error) {

}

*/

////////////////////////////////////////////////////////// Funciones varias//////////////////////////////////////////////

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

func Abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendRabbit(player int32, round int32) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
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
