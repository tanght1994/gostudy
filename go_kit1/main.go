package main

import (
	"net/http"
	"tanght/endpoint"
	"tanght/server"
	"tanght/transport"

	kittransport "github.com/go-kit/kit/transport/http"
)

// 服务发布

func main() {
	s := server.Server{}
	hello := endpoint.MakeServerEndPointHello(s)
	Bye := endpoint.MakeServerEndPointBye(s)
	helloServer := kittransport.NewServer(hello, transport.HelloDecodeRequest, transport.HelloEncodeResponse)
	sayServer := kittransport.NewServer(Bye, transport.ByeDecodeRequest, transport.ByeEncodeResponse)

	// 使用http包启动服务
	go http.ListenAndServe("0.0.0.0:8000", helloServer)

	go http.ListenAndServe("0.0.0.0:8001", sayServer)
	select {}

}
