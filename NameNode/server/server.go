package main

import (
	"bufio"
	"context"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fvalladaresj/SD-Tarea2/Jugador/apiJugador"
	"github.com/fvalladaresj/SD-Tarea2/Lider/api"
	"github.com/fvalladaresj/SD-Tarea2/NameNode/apiNameNode"
	"github.com/fvalladaresj/SD-Tarea2/Pozo/apiPozo"

	"google.golang.org/grpc"
)

type server struct {
	apiNameNode.UnimplementedNameNodeServer
}

// 						   lider		   jugador		     pozo
var ports = [3]string{"0.0.0.0:50051", "0.0.0.0:50053", "0.0.0.0:50054"}

// 						   api		     apiJugador		     apiPozo

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 50052
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := grpc.NewServer()
	// attach the Lider service to the server
	apiNameNode.RegisterNameNodeServer(s, &server{})
	// start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func (*server) EscribirJugada(ctx context.Context, in *apiNameNode.JugadaJugador) (*apiNameNode.Signal, error) {

	var str_Idjugador string = strconv.FormatInt(int64(in.IdJugador), 10)
	var str_Etapa string = strconv.FormatInt(int64(in.Etapa), 10)

	rand.Seed(time.Now().UnixNano())
	port := rand.Int31n(3)

	//Escribir en NameNode
	fileName := "NameNode.txt"
	toWrite := "Jugador_" + str_Idjugador + " Ronda_" + str_Etapa + " " + ports[port]

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := f.WriteString(toWrite + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	//Escribir en DataNode Correspondiente
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(ports[port], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	// 						   lider		   jugador		     pozo
	//var ports = [3]string{"0.0.0.0:50051", "0.0.0.0:50053", "0.0.0.0:50054"}
	// 						   api		     apiJugador		     apiPozo

	if port == 0 {
		c := api.NewLiderClient(conn)
		c.EscribirJugada(context.Background(), &api.JugadaJugador{IdJugador: in.IdJugador, Jugada: in.Jugada, Etapa: in.Etapa})
	} else if port == 1 {
		c := apiJugador.NewDataNodeJugadorClient(conn)
		c.EscribirJugada(context.Background(), &apiJugador.JugadaJugador{IdJugador: in.IdJugador, Jugada: in.Jugada, Etapa: in.Etapa})
	} else if port == 2 {
		c := apiPozo.NewDataNodePozoClient(conn)
		c.EscribirJugada(context.Background(), &apiPozo.JugadaJugador{IdJugador: in.IdJugador, Jugada: in.Jugada, Etapa: in.Etapa})
	}

	return &apiNameNode.Signal{Sign: true}, nil
}

func (*server) PedirJugadasJugador(ctx context.Context, in *apiNameNode.Jugador) (*apiNameNode.TodasLasJugadas, error) {
	str_Idjugador := strconv.FormatInt(int64(in.IdJugador), 10)
	etapa1 := "Jugador_" + str_Idjugador + " Ronda_1"
	etapa2 := "Jugador_" + str_Idjugador + " Ronda_2"
	etapa3 := "Jugador_" + str_Idjugador + " Ronda_3"
	jugadas := ""

	f, err := os.Open("NameNode.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if strings.Join(line[0:2], " ") == etapa1 {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(line[2], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()
			if line[2] == ports[0] {
				c := api.NewLiderClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &api.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(1)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 1: " + response.JugadasJugador
			} else if line[2] == ports[1] {
				c := apiJugador.NewDataNodeJugadorClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiJugador.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(1)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 1: " + response.JugadasJugador
			} else if line[2] == ports[2] {
				c := apiPozo.NewDataNodePozoClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiPozo.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(1)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 1: " + response.JugadasJugador
			}
		} else if strings.Join(line[0:2], " ") == etapa2 {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(line[2], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()
			if line[2] == ports[0] {
				c := api.NewLiderClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &api.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(2)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 2: " + response.JugadasJugador
			} else if line[2] == ports[1] {
				c := apiJugador.NewDataNodeJugadorClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiJugador.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(2)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 2: " + response.JugadasJugador
			} else if line[2] == ports[2] {
				c := apiPozo.NewDataNodePozoClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiPozo.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(2)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 2: " + response.JugadasJugador
			}
		} else if strings.Join(line[0:2], " ") == etapa3 {
			var conn *grpc.ClientConn
			conn, err := grpc.Dial(line[2], grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %s", err)
			}
			defer conn.Close()
			if line[2] == ports[0] {
				c := api.NewLiderClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &api.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(3)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 3: " + response.JugadasJugador
			} else if line[2] == ports[1] {
				c := apiJugador.NewDataNodeJugadorClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiJugador.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(3)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 3: " + response.JugadasJugador
			} else if line[2] == ports[2] {
				c := apiPozo.NewDataNodePozoClient(conn)
				response, err := c.RetornarJugadas(context.Background(), &apiPozo.JugadorYEtapa{IdJugador: in.IdJugador, NroEtapa: int32(3)})
				if err != nil {
					log.Fatal(err)
				}
				jugadas = jugadas + "Jugadas ronda 3: " + response.JugadasJugador
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &apiNameNode.TodasLasJugadas{JugadasJugador: jugadas}, nil
}
