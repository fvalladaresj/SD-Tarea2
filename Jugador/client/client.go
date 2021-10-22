package client

import(
	"fmt"
	"https://github.com/fvalladaresj/SD-Tarea2/tree/main/Lider/api"
	"golang.org/x/net/context"
  	"google.golang.org/grpc"
	"math/rand"
)

func doPlay(etapa int, playerMove int ) [16]int {
	if etapa == 1 {
		var result [16]int
		result[0] = playerMove 
		for i:=1; i<16; i++ {
			result[i] = rand.Intn(9)+1
		}
	}
	return result
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
	  log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := api.NewLiderClient(conn)


	var input string
	
	fmt.Println("Para participar en el Juego del calamar ingrese \'participar\' ")
	for {
		fmt.Scanln(&input)
		response, err = c.ParticiparJuego(context.Background(), &api.PeticionParticipar{participar = input})
		if !(response != "ya esta participando en el juego del calamar"){
			fmt.Println(response)
			break
		}
		else{
			fmt.Println(response)
		}
	}
	fmt.Println("Espere a las instrucciones para comenzar primera etapa")
	for {
		response, err = c.EstadoEtapas(context.Background(), &api.Check{signal = 1})
		if !(response == 0){
			break
		}
		else{
			fmt.Println("Ha comenzado la primera etapa, Luz Verde, Luz Roja, ingrese un numero entre el 1 y el 10")
		}
	}

	response, err = c.Jugar(context.Background(), &api.Jugadas{jugadas: })



	

	if 
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
	  log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}