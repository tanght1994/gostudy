package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var NoDataTimeOut time.Duration

func main() {
	port := flag.Int("port", 9999, "set the tcp port")
	ndto := flag.Int("ndto", 1, "no data time out (s)")
	flag.Parse()
	NoDataTimeOut = time.Duration(*ndto) * time.Second
	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	li, err := net.Listen("tcp", addr)
	if err != nil {
		exit(err.Error())
	}
	fmt.Printf("Listen on %s\n", addr)

	for {
		cn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(cn)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func handle(cn net.Conn) {
	defer cn.Close()
	local := cn.LocalAddr().String()
	remote := cn.RemoteAddr().String()
	stop := time.Now().Add(NoDataTimeOut)
	recv := []byte{}
	buf := make([]byte, 4096)
	for time.Now().Before(stop) {
		cn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
		n, err := cn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			if operr, ok := err.(*net.OpError); ok && operr.Timeout() {
				continue
			}
			fmt.Printf("Read error: %v\n", err)
			return
		}
		if n != 0 {
			recv = append(recv, buf[:n]...)
			stop = time.Now().Add(NoDataTimeOut)
		}
	}
	send := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain; charset=utf-8\r\n\r\nmyaddr:%s\r\nyouraddr:%s\r\n%s\r\n\r\n", local, remote, recv)
	// Date: Thu, 05 May 2022 17:06:07 GMT\r\n
	_, err := cn.Write([]byte(send))
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	}
}
