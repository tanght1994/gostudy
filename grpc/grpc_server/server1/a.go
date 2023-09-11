package server1

import (
	"context"
	"tanght/proto/pbs1"
)

// S1 ...
type S1 struct {
	pbs1.UnimplementedS1Server
}

// Hello ...
func (s1 S1) Hello(context.Context, *pbs1.HelloReq) (*pbs1.HelloRes, error) {
	return &pbs1.HelloRes{Msg: "S1 Hello"}, nil
}

// Hi ...
func (s1 S1) Hi(context.Context, *pbs1.HiReq) (*pbs1.HiRes, error) {
	return &pbs1.HiRes{Msg: "S1 Hi"}, nil
}
