package client

import(
	"fmt"
	"https://github.com/fvalladaresj/SD-Tarea2/tree/main/Lider/api"
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
	fmt.Println("Para participar en el Juego del calamar ingrese \'participar\' ")
	for {
		fmt.Scanln(&input)
		response, err = c.ParticiparJuego(context.Background(), &api.PeticionParticipar{participar = input})
		if !(response == "ya esta participando en el juego del calamar"){
			fmt.Println(response)
			break
		}
	for {
		response, err = c.Estado(context.Background(), &api.Check{signal = 1})
		if !(response == "ya esta participando en el juego del calamar"){
			fmt.Println(response)
			break
		}
	}

	


	
	
	

	if 
	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
	  log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
  }