package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10001")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		go handle_conn(conn)
	}
}

func handle_conn(conn net.Conn) {
	defer conn.Close()
	data := []byte{}
	tmp := make([]byte, 4096)
	t0 := time.Now()
	for {
		if time.Since(t0) > 1*time.Second {
			break
		}
		err := conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		if err != nil {
			break
		}
		n, err := conn.Read(tmp)
		if err == io.EOF {
			break
		}
		if err, ok := err.(*net.OpError); ok && err.Timeout() {
			continue
		}
		if n == 0 {
			continue
		}
		data = append(data, tmp[0:n]...)
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(data))
	conn.Write([]byte("\r\n"))
	conn.Write(data)
	conn.Write([]byte("\r\n\r\n"))
}
