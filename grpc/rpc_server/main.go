package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type RPC_Rect struct{}

type RPC_Params struct {
	W, H int
}

func (r *RPC_Rect) Area(p RPC_Params, ret *int) error {
	*ret = p.H * p.W
	return nil
}

func (r *RPC_Rect) Perimeter(p RPC_Params, ret *int) error {
	*ret = (p.H + p.W) * 2
	return nil
}

func rpc_server() {
	rect := new(RPC_Rect)
	rpc.Register(rect)
	rpc.HandleHTTP()
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		fmt.Println("server error: ", err)
	}
}

func main() {
	rpc_server()
}
