package main

import (
	"context"
	"fmt"
	"log"

	"github.com/matiniiuu/go-grpc/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm Client...")
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Faild to connect: %v", err)
	}
	defer connection.Close()
	client := greetpb.NewGreetServiceClient(connection)
	doUnary(client)

}

func doUnary(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "matin",
			LastName:  "matiniiuu",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Calling rpc: %v", err)
	}
	log.Printf("Response : %v\n", res.Result)
}
