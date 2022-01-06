package main

import (
	"context"
	message "tanght/protobuf"

	"google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	message.RegisterMessageSenderServer(srv, &MessageSenderServer{})

}

type MessageSenderServer struct {
	*message.UnimplementedMessageSenderServer
}

func (m *MessageSenderServer) Send(context.Context, *message.MessageRequest) (*message.MessageResponse, error) {
	return nil, nil
}
