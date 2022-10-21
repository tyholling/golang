// Command server
package main

import (
	"net"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	err := os.MkdirAll("log", 0o0644)
	if err != nil {
		log.Fatalf("failed to create log directory: %s", err)
	}
	file, err := os.Create("log/server.log")
	if err != nil {
		log.Fatalf("failed to create log file: %s", err)
	} else {
		log.SetOutput(file)
	}
}

func main() {
	log.Info("server: started")
	defer log.Info("server: stopped")

	conn, err := net.Listen("tcp", "localhost:65000")
	if err != nil {
		log.Errorf("failed to listen: %s", err)
		return
	}

	server := grpc.NewServer()
	pb.RegisterConnectionServiceServer(server, &connectionServer{})
	err = server.Serve(conn)
	if err != nil {
		log.Errorf("failed to start server: %s", err)
		return
	}
}

type connectionServer struct {
	pb.UnimplementedConnectionServiceServer
}

func (s *connectionServer) Connect(stream pb.ConnectionService_ConnectServer) error {
	wg := sync.WaitGroup{}
	messageChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range messageChan {
			msg := &pb.ConnectResponse{}
			err := stream.Send(msg)
			if err != nil {
				log.Errorf("failed to send: %s", err)
				continue
			}
			log.Debugf("send: %s", msg)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Errorf("failed to receive: %s", err)
				continue
			}
			messageChan <- struct{}{}
			log.Debugf("receive: %s", msg)
		}
	}()

	wg.Wait()
	return nil
}
