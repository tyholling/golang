package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.ConnectionServiceClient
	stream pb.ConnectionService_ConnectClient
}

func (c *Client) Connect() error {
	target := "localhost:65000"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect to target: %w", err)
	}
	c.conn = conn

	c.client = pb.NewConnectionServiceClient(c.conn)
	stream, err := c.client.Connect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to connect to server: %w", err)
	}
	c.stream = stream

	return nil
}

func (c *Client) Start() {
	wg := sync.WaitGroup{}
	messages := make(chan *pb.ConnectRequest)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.handleSend(messages)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.handleRecv(messages)
	}()

	wg.Wait()
}

func (c *Client) handleSend(messages <-chan *pb.ConnectRequest) {
	for msg := range messages {
		err := c.stream.Send(msg)
		if err != nil {
			log.Print("failed to send: ", err)
			continue
		}
		log.Printf("send: %s", msg)
	}
}

func (c *Client) handleRecv(_ chan<- *pb.ConnectRequest) {
	for {
		msg, err := c.stream.Recv()
		if err != nil {
			log.Print("failed to receive: ", err)
			c.reconnect()
			continue
		}

		if msg.Request != nil {
			request, err := anypb.UnmarshalNew(msg.Request, proto.UnmarshalOptions{})
			if err != nil {
				log.Print(err)
				continue
			}
			log.Printf("received request: %s", request)
		} else if msg.Response != nil {
			response, err := anypb.UnmarshalNew(msg.Response, proto.UnmarshalOptions{})
			if err != nil {
				log.Print(err)
				continue
			}
			log.Printf("received response: %s", response)
		}
	}
}

func (c *Client) reconnect() {
	for {
		stream, err := c.client.Connect(context.Background())
		if err == nil {
			c.stream = stream
			log.Print("reconnected to server")
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	client := &Client{}
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
