package main

import (
	"context"
	"fmt"
	"os"
	proto "test/proto"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

type Hello struct {
}

func (g *Hello) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func runClient(service micro.Service) {
	// Create new greeter client
	greeter := proto.NewHelloService("hello", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Greeting)
}

func main() {
	service := micro.NewService(
		micro.Name("hello"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloword",
		}),
		micro.Flags(cli.BoolFlag{
			Name:  "run_client",
			Usage: "Launch the client",
		}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)
	proto.RegisterHelloHandler(service.Server(), new(Hello))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
