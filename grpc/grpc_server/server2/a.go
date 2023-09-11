package server2

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
	"tanght/proto/pbs2"
	"time"
)

// S2 ...
type S2 struct {
	pbs2.UnimplementedS2Server
}

// ClientStream 客户端流
// Recv() EOF SendAndClose
func (s S2) ClientStream(stream pbs2.S2_ClientStreamServer) error {
	reqList := []string{}
	for {
		req, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				log.Printf("recv from client error %v\n", err)
			}
			break
		}
		reqList = append(reqList, req.Msg)
	}
	msg := "i have get your req, your req is " + strings.Join(reqList, ",")
	stream.SendAndClose(&pbs2.ClientStreamRes{Msg: msg})
	return nil
}

// ServerStream 服务端流
// Send()
func (s S2) ServerStream(req *pbs2.ServerStreamReq, stream pbs2.S2_ServerStreamServer) error {
	for i := 0; i < 5; i++ {
		stream.Send(&pbs2.ServerStreamRes{Msg: fmt.Sprintf("haha%d", i)})
		time.Sleep(time.Second)
	}
	// client端有stream.CloseSend() server端没有？
	// 此函数返回时, server的stream自动关闭, 对应client端的stream.CloseSend()
	return nil
}

// TwoWayStream 双向流
func (s S2) TwoWayStream(stream pbs2.S2_TwoWayStreamServer) error {
	log.Printf("TwoWayStream start\n")
	defer log.Printf("TwoWayStream end\n")
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err != nil {
				if err != io.EOF {
					log.Printf("recv from client error %v\n", err)
				}
				log.Printf("recv from client over\n")
				break
			}
			log.Printf("recv from client %s\n", req.Msg)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			stream.Send(&pbs2.TwoWayStreamRes{Msg: strconv.Itoa(i)})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	// client端有stream.CloseSend() server端没有？
	// 此函数返回时, server的stream自动关闭, 对应client端的stream.CloseSend()
	return nil
}
