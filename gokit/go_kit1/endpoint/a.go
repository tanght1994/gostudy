package endpoint

import (
	"context"
	"tanght/server"

	"github.com/go-kit/kit/endpoint"
)

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Reply string `json:"reply"`
}

func MakeServerEndPointHello(s server.IServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(HelloRequest)
		if !ok {
			return HelloResponse{}, nil
		}
		return HelloResponse{Reply: s.Hello(r.Name)}, nil
	}
}

type ByeRequest struct {
	Name string `json:"name"`
}

type ByeResponse struct {
	Reply string `json:"reply"`
}

func MakeServerEndPointBye(s server.IServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(ByeRequest)
		if !ok {
			return ByeResponse{}, nil
		}
		return ByeResponse{Reply: s.Bye(r.Name)}, nil
	}
}
