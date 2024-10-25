package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	chat "grpcchat/chat_pb"
	"log"
	"os"
	"time"
)

func writeRoutine(end chan interface{}, ctx context.Context, conn chat.ChatExampleClient) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			if str == "exit" {
				break OUTER
			}
			msg, err := conn.SendMessage(context.Background(),
				&chat.ChatMessage{Text: str})
			if err != nil {
				fmt.Printf("error %s", status.Convert(err).Message())
			}
			if msg != nil {
				created, err := ptypes.Timestamp(msg.Created)
				if err != nil {
					fmt.Printf("error %s", err)
				}
				fmt.Printf("[%s] id: %d msg: %s", created.Format(time.RFC3339Nano), msg.Id, msg.Text)
			}

		}
	}
	close(end)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := chat.NewChatExampleClient(cc)
	end := make(chan interface{})
	go writeRoutine(end, ctx, c)
	<-end
}
