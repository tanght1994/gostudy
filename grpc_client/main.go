package main

import (
	"context"
	"fmt"
	"log"
	"tanght/pbhello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	clie := pbhello.NewHelloWorldClient(conn)
	ctx := context.Background()
	req := &pbhello.HelloReq{Name: "tanght"}
	res, err := clie.Hello(ctx, req)
	if err != nil {
		log.Fatalf("clie.Hello error: %v", err)
	}
	fmt.Println(res.Msg)
}
