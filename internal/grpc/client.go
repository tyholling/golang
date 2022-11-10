// Package grpc handles the grpc connection
package grpc

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
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
					go handleHeartbeat(msgChan)
				case pb.SubscriptionType_METRICS:
					go handleMetrics(msgChan)
				}
			}
		} else if msg.Response != nil {
			log.Debugf("received response: %s", msg)
		}
	}
}

func handleHeartbeat(msgChan chan<- *pb.ConnectRequest) {
	ticker := time.NewTicker(time.Minute)
	for {
		response, err := anypb.New(&pb.Heartbeat{
			Timestamp: timestamppb.Now(),
		})
		if err != nil {
			log.Error(err)
			continue
		}
		msgChan <- &pb.ConnectRequest{
			Response: response,
		}

		<-ticker.C
	}
}

func handleMetrics(msgChan chan<- *pb.ConnectRequest) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		metrics := &pb.Metrics{}

		pz, err := cpu.Percent(time.Second, false)
		if err != nil {
			log.Error(err)
			continue
		}
		if len(pz) == 0 {
			log.Error("failed to retrieve CPU utilization")
			continue
		}
		metrics.Cpu = pz[0]

		mz, err := mem.VirtualMemory()
		if err != nil {
			log.Error(err)
			continue
		}
		if mz == nil {
			log.Error("failed to retrieve memory utilization")
			continue
		}
		metrics.Memory = mz.UsedPercent

		izz, err := net.IOCounters(false)
		if err != nil {
			log.Error(err)
			continue
		}
		if len(izz) == 0 {
			log.Error("failed to retrieve network metrics")
			continue
		}
		iz := izz[0]
		metrics.BytesSent = iz.BytesSent
		metrics.BytesReceived = iz.BytesRecv
		metrics.ErrorsIn = iz.Errin
		metrics.ErrorsOut = iz.Errout
		metrics.DiscardsIn = iz.Dropin
		metrics.DiscardsOut = iz.Dropout

		response, err := anypb.New(metrics)
		if err != nil {
			log.Error(err)
			continue
		}
		msgChan <- &pb.ConnectRequest{
			Response: response,
		}

		<-ticker.C
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
