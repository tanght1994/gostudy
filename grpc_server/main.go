package main

import (
	"context"
	"fmt"
	"net"
	"tanght/grpchello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	grpchello.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	err = s.Serve(lis)
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *grpchello.ReqHello) (*grpchello.RepHello, error) {
	return &grpchello.RepHello{Msg: "ni hao a " + in.Name}, nil
}
