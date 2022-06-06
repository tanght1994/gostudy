package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:10001")
	if err != nil {
		return
	}
	fmt.Println(conn)
	time.Sleep(100 * time.Hour)
}
