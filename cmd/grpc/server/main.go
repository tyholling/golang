package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

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

	wg.Wait()
	return nil
}

func handleSend(stream pb.ConnectionService_ConnectServer, messages <-chan *pb.ConnectResponse) {
	for msg := range messages {
		err := stream.Send(msg)
		if err != nil {
			log.Fatalf("failed to send: %s", err)
		}
		log.Printf("send: %s", msg)
	}
}

func handleRecv(stream pb.ConnectionService_ConnectServer, _ chan<- *pb.ConnectResponse) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive: %s", err)
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

func main() {
	server := &Server{}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
