package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/fvalladaresj/SD-Tarea2/Pozo/apiPozo"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var pool = 0

type server struct {
	apiPozo.UnimplementedDataNodePozoServer
}

func main() {
	// create a listener on TCP port 50054
	lis, err := net.Listen("tcp", "0.0.0.0:50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := grpc.NewServer()
	// attach the Lider service to the server
	apiPozo.RegisterDataNodePozoServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (*server) EscribirJugada(ctx context.Context, in *apiPozo.JugadaJugador) (*apiPozo.Signal, error) {

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

	return &apiPozo.Signal{Sign: true}, nil

}

func (*server) RetornarJugadas(ctx context.Context, in *apiPozo.JugadorYEtapa) (*apiPozo.JugadasArchivo, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_NroEtapa string = strconv.FormatInt(int64(in.NroEtapa), 10)

	var nombre_archivo string = "jugador_" + str_Idjugador + "__ronda" + str_NroEtapa + ".txt"

	content, err := os.ReadFile(nombre_archivo)
	if err != nil {
		log.Fatal(err)
	}

	var string_content string = string(content)

	return &apiPozo.JugadasArchivo{JugadasJugador: string_content}, nil

}

func listenRabbit() {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			pool += 100000000
			EscribirMuerto(string(d.Body), pool)
		}
	}()

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func EscribirMuerto(deathInfo string, pool int) {
	strPool := strconv.FormatInt(int64(pool), 10)

	toWrite := deathInfo + " " + strPool

	f, err := os.OpenFile("Pool.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(toWrite + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}
}
