package main

import (
	"fmt"
	"tanght/protobuf/pba"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("123")
	p := &pba.Person{}
	proto.Marshal(p)
	proto.Unmarshal()
}
protoc --go_out=plugins=grpc:. xxxx.proto