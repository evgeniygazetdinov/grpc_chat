package main

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	chat "grpcchat/chat_pb"
	"log"
	"net"
)

type ServerExample struct {
}

var i int64

func (s ServerExample) SendMessage(ctx context.Context, msg *chat.ChatMessage) (*chat.ChatMessage, error) {
	defer func() { i++ }() // странный икремент для получения айди
	if msg.Text == "" {
		return nil, status.Error(codes.InvalidArgument, "No empty allowed")
	}
	return &chat.ChatMessage{Text: "PONG " + msg.Text, Id: i, Created: ptypes.TimestampNow()}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat.RegisterChatServiceServer(grpcServer, ServerExample{})
	_ = grpcServer.Serve(lis)

}
