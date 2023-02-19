package main

import (
	pb "greetdemo/greet/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

const (
	addr string = "0.0.0.0:8080"
)

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("listening at %s", addr)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}
