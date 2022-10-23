// Package grpc handles the grpc connection
package grpc

import (
	"fmt"
	"net"
	"sync"

	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Server represents the grpc server
type Server struct {
	conn   net.Listener
	server *grpc.Server
}

// Listen listens for new grpc connections
func (s *Server) Listen() error {
	conn, err := net.Listen("tcp", "localhost:65000")
	if err != nil {
		return fmt.Errorf("failed to listen: %s", err)
	}
	s.conn = conn

	return nil
}

// Start accepts incoming grpc connections
func (s *Server) Start() {
	s.server = grpc.NewServer()
	pb.RegisterConnectionServiceServer(s.server, &connectionServer{})
	err := s.server.Serve(s.conn)
	if err != nil {
		log.Errorf("failed to start server: %s", err)
		return
	}
}

type connectionServer struct {
	pb.UnimplementedConnectionServiceServer
}

// Connect runs the handlers to send and receive grpc messages
func (s *connectionServer) Connect(stream pb.ConnectionService_ConnectServer) error {
	wg := sync.WaitGroup{}
	msgChan := make(chan *pb.ConnectResponse)

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleSend(stream, msgChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleRecv(stream, msgChan)
	}()

	wg.Wait()
	return nil
}

func handleSend(stream pb.ConnectionService_ConnectServer, msgChan <-chan *pb.ConnectResponse) {
	for msg := range msgChan {
		err := stream.Send(msg)
		if err != nil {
			log.Errorf("failed to send: %s", err)
			continue
		}
		log.Debugf("send: %s", msg)
	}
}

func handleRecv(stream pb.ConnectionService_ConnectServer, msgChan chan<- *pb.ConnectResponse) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Errorf("failed to receive: %s", err)
			continue
		}

		if msg.Request != nil {
			log.Debugf("received request: %s", msg)

			request, err := anypb.UnmarshalNew(msg.Request, proto.UnmarshalOptions{})
			if err != nil {
				log.Error(err)
				continue
			}
			if v, ok := request.(*pb.PingRequest); ok {
				response, err := anypb.New(&pb.PingResponse{
					Timestamp: v.Timestamp,
				})
				if err != nil {
					log.Error(err)
					continue
				}

				msg := &pb.ConnectResponse{
					Response: response,
				}
				msgChan <- msg
			}
		} else if msg.Response != nil {
			log.Debugf("received response: %s", msg)
		}
	}
}
