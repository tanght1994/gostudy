package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务器
	c, e := net.Dial("tcp", "127.0.0.1:8000")
	if e != nil {
		fmt.Printf("Dial error， %v\n", e)
		return
	}
	fmt.Printf("Dial %s %s %s %s\n", c.LocalAddr().Network(), c.LocalAddr().String(), c.RemoteAddr().Network(), c.RemoteAddr().String())

	// 写消息
	send := "hello world"
	_, e = c.Write([]byte(send))
	if e != nil {
		fmt.Printf("Write error， %v\n", e)
		return
	}
	fmt.Printf("Write %s\n", send)

	//读消息
	buff := make([]byte, 1024)
	n, e := c.Read(buff)
	if e != nil {
		fmt.Printf("Read error， %v\n", e)
		return
	}
	recv := string(buff[0:n])
	fmt.Printf("Read %s\n", recv)

	// 关闭连接
	e = c.Close()
	if e != nil {
		fmt.Printf("Close error， %v\n", e)
		return
	}
	fmt.Printf("Close\n")
}
