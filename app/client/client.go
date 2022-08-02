package main

import (
	"context"
	"io"
	"log"
	"message-delivery-system/proto/pb"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

const (
	address  = "localhost:8080"
)

func main() {
	messageType := os.Args[1]

	var message string
	users := make([]uint64, 0)
	var err error

	if messageType == "relay" {
		userId := os.Args[2]
		message = os.Args[3]

		s := strings.Split(userId, ",")
		for _, i := range s {
			user, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				log.Fatalf("failed to convert userid: %v", err)
			}
			users = append(users, user)
		}

	}

	log.Printf("Message Type: %v, Content: %v", messageType, message)

	conn, err := grpc.Dial(
		address,
	)

	if err != nil {
		log.Fatalf("failed to Serve: %v", err)
	}

	client := pb.NewDeliveryServiceClient(conn)

	// create stream
	stream, err := client.Delivery(context.Background())
	if err != nil {
		log.Println(err)
	}
	// Send Message to stream
	req := pb.Request{Message: message, Type: messageType, Users: users}
	if err := stream.Send(&req); err != nil {
		log.Println(err)
	}

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Fatalf("can not receive %v", err)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("received: %v", resp)
		}
	}()

	time.Sleep(5 * time.Minute)
}
