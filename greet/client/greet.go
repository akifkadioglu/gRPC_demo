package main

import (
	"context"
	"fmt"
	pb "greetdemo/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "akif",
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf(res.Result)
}
