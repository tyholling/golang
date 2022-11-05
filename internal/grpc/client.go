// Package grpc handles the grpc connection
package grpc

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
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
			c.reconnect()
			continue
		}

		if msg.Request != nil {
			log.Debugf("received request: %s", msg)

			request, err := anypb.UnmarshalNew(msg.Request, proto.UnmarshalOptions{})
			if err != nil {
				log.Error(err)
				continue
			}
			if v, ok := request.(*pb.Subscribe); ok {
				switch v.Type {
				case pb.SubscriptionType_HEARTBEAT:
					go func() {
						ticker := time.NewTicker(time.Second)
						for {
							<-ticker.C

							response, err := anypb.New(&pb.Heartbeat{
								Timestamp: timestamppb.Now(),
							})
							if err != nil {
								log.Error(err)
								continue
							}
							msg := &pb.ConnectRequest{
								Response: response,
							}
							msgChan <- msg
						}
					}()
				case pb.SubscriptionType_CPU:
					go func() {
						ticker := time.NewTicker(time.Second)
						for {
							<-ticker.C

							pz, err := cpu.Percent(time.Minute, false)
							if err != nil {
								log.Error(err)
								continue
							}
							if len(pz) == 0 {
								log.Error("failed to retrieve CPU utilization")
								continue
							}

							response, err := anypb.New(&pb.MetricsCPU{
								Percent: pz[0],
							})
							if err != nil {
								log.Error(err)
								continue
							}
							msg := &pb.ConnectRequest{
								Response: response,
							}
							msgChan <- msg
						}
					}()
				}
			}
		} else if msg.Response != nil {
			log.Debugf("received response: %s", msg)
		}
	}
}

func (c *Client) reconnect() {
	for {
		stream, err := c.client.Connect(context.Background())
		if err == nil {
			c.stream = stream
			log.Info("reconnected to server")
			break
		}
		time.Sleep(time.Second)
	}
}
