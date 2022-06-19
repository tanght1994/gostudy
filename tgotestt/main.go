package main

import (
	"context"
	"net"
	"tanght/fuckproto"

	"google.golang.org/grpc"
)

type GreeterServer struct {
	fuckproto.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *fuckproto.HelloRequest) (*fuckproto.HelloReply, error) {
	return nil, nil
}
func (s *GreeterServer) SayHelloAgain(context.Context, *fuckproto.HelloRequest) (*fuckproto.HelloReply, error) {
	return nil, nil
}

func (s *GreeterServer) SayHello1(ss fuckproto.Greeter_SayHello1Server) error {
	ss.Recv()
	return nil
}
func (s *GreeterServer) SayHello2(req *fuckproto.HelloRequest, ss fuckproto.Greeter_SayHello2Server) error {
	return nil
}
func (s *GreeterServer) SayHello3(ss fuckproto.Greeter_SayHello3Server) error {
	return nil
}

func newGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

func main() {
	root := grpc.NewServer()
	greeterServer := newGreeterServer()
	fuckproto.RegisterGreeterServer(root, greeterServer)
	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	must(err)
	root.Serve(lis)
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
