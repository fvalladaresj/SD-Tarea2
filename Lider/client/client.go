package client

import(
	"fmt"
	"https://github.com/fvalladaresj/SD-Tarea2/tree/main/Lider/api"
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
		if !(input == "participar"){
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