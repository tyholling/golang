package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)

	server := &Server{}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type Server struct {
	listener net.Listener
	server   *grpc.Server
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", "localhost:65000")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.listener = listener

	s.server = grpc.NewServer()
	pb.RegisterConnectionServiceServer(s.server, &connectionServer{})

	err = s.server.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

type connectionServer struct {
	pb.UnimplementedConnectionServiceServer
}

func (s *connectionServer) Connect(stream pb.ConnectionService_ConnectServer) error {
	wg := sync.WaitGroup{}
	messages := make(chan *pb.ConnectResponse)

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleSend(stream, messages)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleRecv(stream, messages)
	}()

	request, err := anypb.New(&pb.Subscribe{
		Type: pb.Subscription_SUBSCRIPTION_HEARTBEAT,
	})
	if err != nil {
		log.Print(err)
	}
	messages <- &pb.ConnectResponse{
		Request: request,
	}

	wg.Wait()
	return nil
}

func handleSend(stream pb.ConnectionService_ConnectServer, messages <-chan *pb.ConnectResponse) {
	for msg := range messages {
		err := stream.Send(msg)
		if err != nil {
			log.Printf("failed to send: %s", err)
			return
		}
		log.Printf("send: %s", msg)
	}
}

func handleRecv(stream pb.ConnectionService_ConnectServer, _ chan<- *pb.ConnectResponse) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Printf("failed to receive: %s", err)
			return
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
