package main

import (
	"grpctest/echo"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoServer struct{}

func (s *EchoServer) Echo(ctx context.Context, in *echo.EchoMessage) (*echo.EchoMessage, error) {
	return &echo.EchoMessage{Msg: "Hello, World !!!!"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &EchoServer{})
	s.Serve(l)
}
