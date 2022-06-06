package main

import (
	"context"
	"fmt"
	"net"
	"tanght/pbhello"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	pbhello.RegisterHelloWorldServer(s, helloWorldServer{})
	s.Serve(lis)
}

type helloWorldServer struct {
	pbhello.UnimplementedHelloWorldServer
}

func (helloWorldServer) Hello(ctx context.Context, req *pbhello.HelloReq) (*pbhello.HelloRes, error) {
	fmt.Println("in Hello")
	return &pbhello.HelloRes{Msg: fmt.Sprint("Hello ", req.Name)}, nil
}
func (helloWorldServer) Hi(ctx context.Context, req *pbhello.HiReq) (*pbhello.HiRes, error) {
	fmt.Println("in Hi")
	return &pbhello.HiRes{Msg: fmt.Sprint("Hi ", req.Name, " your age is ", req.Age)}, nil
}
