package main

import (
	"context"
	"fmt"
	"tanght/grpchello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接服务端失败: %s", err)
		return
	}
	defer conn.Close()

	c := grpchello.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &grpchello.ReqHello{Name: "horika"})
	if err != nil {
		fmt.Printf("调用服务端代码失败: %s", err)
		return
	}
	fmt.Printf("调用成功: %s", r.Msg)
}
