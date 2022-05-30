package main

import (
	"fmt"
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

func rpc_client() {
	c, e := rpc.DialHTTP("tcp", "127.0.0.1:8000")
	if e != nil {
		fmt.Println("client error: ", e)
	}
	r := 0
	e = c.Call("RPC_Rect.Area", RPC_Params{3, 5}, &r)
	if e != nil {
		fmt.Println("client error: ", e)
	}
	fmt.Println("client res: ", r)
	e = c.Call("RPC_Rect.Perimeter", RPC_Params{3, 5}, &r)
	if e != nil {
		fmt.Println("client error: ", e)
	}
	fmt.Println("client res: ", r)
}

func main() {
	rpc_client()
}
