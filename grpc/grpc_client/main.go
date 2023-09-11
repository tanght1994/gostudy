package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"sync"
	"tanght/proto/pbs2"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbs2.NewS2Client(conn)
	// ServerStream(client)
	// ClientStream(client)
	TwoWayStream(client)
}

// ServerStream 服务端流式
func ServerStream(grpcClient pbs2.S2Client) {
	log.Println("ServerStream start")
	defer log.Println("ServerStream end")
	stream, err := grpcClient.ServerStream(context.TODO(), &pbs2.ServerStreamReq{})
	if err != nil {
		return
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Printf("recv from server error %v\n", err)
			}
			break
		}
		log.Printf("res is %s\n", res.Msg)
	}
}

// ClientStream 客户端流式
func ClientStream(grpcClient pbs2.S2Client) {
	log.Println("ClientStream start")
	defer log.Println("ClientStream end")
	stream, err := grpcClient.ClientStream(context.TODO())
	if err != nil {
		return
	}
	for i := 0; i < 5; i++ {
		err = stream.Send(&pbs2.ClientStreamReq{Msg: fmt.Sprintf("%d", i)})
		if err != nil {
			log.Printf("send to server error %v\n", err)
			break
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("recv from server error %v\n", err)
		return
	}
	log.Printf("res is %s\n", res.Msg)
}

// TwoWayStream 双向流式
func TwoWayStream(grpcClient pbs2.S2Client) {
	log.Println("TwoWayStream start")
	defer log.Println("TwoWayStream end")
	stream, err := grpcClient.TwoWayStream(context.TODO())
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			stream.Send(&pbs2.TwoWayStreamReq{Msg: strconv.Itoa(i)})
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	go func() {
		defer wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Printf("recv from server error %v\n", err)
				}
				break
			}
			log.Printf("res is %s\n", res.Msg)
		}
	}()
	wg.Wait()
}
