package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	// 监听地址
	l, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen error")
		return
	}
	fmt.Printf("Listen %s %s\n", l.Addr().Network(), l.Addr().String())

	for {
		// 接收客户端连接
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error")
			return
		}
		fmt.Printf("Accept %s %s %s %s\n", c.LocalAddr().Network(), c.LocalAddr().String(), c.RemoteAddr().Network(), c.RemoteAddr().String())

		// 读取客户端数据
		buff := make([]byte, 1024)
		n, err := c.Read(buff)
		if err != nil {
			fmt.Println("Read error")
			return
		}
		msg := string(buff[0:n])
		fmt.Printf("Read %s\n", msg)

		// 首字母大写，写回去
		tmp := []byte(msg)
		if len(tmp) > 0 {
			up := bytes.ToUpper(tmp[0:1])
			tmp[0] = up[0]
		}
		_, err = c.Write(tmp)
		if err != nil {
			fmt.Println("Write error")
			return
		}
		fmt.Printf("Write %s\n", string(tmp))

		n, err = c.Read(buff)
		if err != nil {
			fmt.Printf("2 Read error, %v\n", err)
			return
		}
		fmt.Printf("2 Read %d\n", n)
	}
}
