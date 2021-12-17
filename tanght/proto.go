package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func proto_main() {
	a := &ProtoTest1{Name: "tanght", Age: 18}
	b, e := proto.Marshal(a)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(b)
	}
	c := &ProtoTest1{Name: "aaaa", Age: 10}
	fmt.Println(c)
	e = proto.Unmarshal(b, c)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(c)
	}
}
