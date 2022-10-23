// Package grpc handles the grpc connection
package grpc

import (
	"context"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Client represents the grpc client
type Client struct {
	conn   *grpc.ClientConn
	client pb.ConnectionServiceClient
	stream pb.ConnectionService_ConnectClient
}

// Connect creates the grpc connection
func (c *Client) Connect() error {
	target := "localhost:65000"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect to channel: %s", err)
	}
	c.conn = conn

	c.client = pb.NewConnectionServiceClient(c.conn)
	stream, err := c.client.Connect(context.Background())
	if err != nil {
		return fmt.Errorf("failed to connect to server: %s", err)
	}
	c.stream = stream

	return nil
}

// Start runs the handlers to send and receive grpc messages
func (c *Client) Start() {
	wg := sync.WaitGroup{}
	msgChan := make(chan *pb.ConnectRequest)

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.handleSend(msgChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.handleRecv(msgChan)
	}()

	// send ping request
	request, err := anypb.New(&pb.PingRequest{
		Timestamp: timestamppb.Now(),
	})
	if err != nil {
		log.Fatal(err)
	}
	msg := &pb.ConnectRequest{
		Request: request,
	}
	msgChan <- msg

	wg.Wait()
}

func (c *Client) handleSend(msgChan <-chan *pb.ConnectRequest) {
	for msg := range msgChan {
		err := c.stream.Send(msg)
		if err != nil {
			log.Errorf("failed to send: %s", err)
			continue
		}
		log.Debugf("send: %s", msg)
	}
}

func (c *Client) handleRecv(msgChan chan<- *pb.ConnectRequest) {
	for {
		msg, err := c.stream.Recv()
		if err != nil {
			log.Errorf("failed to receive: %s", err)
			continue
		}

		if msg.Request != nil {
			log.Debugf("received request: %s", msg)
		} else if msg.Response != nil {
			log.Debugf("received response: %s", msg)
		}
	}
}
