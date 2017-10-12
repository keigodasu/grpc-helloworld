package main

import (
	fmt "fmt"
	"grpctest/echo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := "localhost:8088"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := echo.NewEchoServiceClient(conn)
	feature, err := client.Echo(context.Background(), &echo.EchoMessage{Msg: "yeeeee"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(feature)
}
