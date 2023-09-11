package main

import (
	"fmt"
	"net"
	"tanght/proto/pbs1"
	"tanght/proto/pbs2"
	"tanght/server1"
	"tanght/server2"

	"google.golang.org/grpc"
)

func main() {
	// var stdout, stderr bytes.Buffer
	// cmd := exec.Command(`C:\Users\tanght\Desktop\tmp\test\go_main\tanght.exe`)
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// cmd.Start()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("+++++++++++")
	// 	fmt.Println(stdout.String())
	// 	fmt.Println("-----------")
	// 	fmt.Println(stderr.String())
	// 	fmt.Println("===========")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Process.Signal(syscall.SIGKILL) %v\n", cmd.Process.Signal(syscall.SIGKILL))
	// fmt.Printf("cmd.Wait(), %v", cmd.Wait())
	// return
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	s := grpc.NewServer()
	pbs1.RegisterS1Server(s, server1.S1{})
	pbs2.RegisterS2Server(s, server2.S2{})
	s.Serve(lis)
}
