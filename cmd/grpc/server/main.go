// command server
package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}

	err := os.MkdirAll("log", 0o0644)
	if err != nil {
		log.Fatalf("failed to create log directory: %s", err)
	}
	file, err := os.Create("log/server.log")
	if err != nil {
		log.Fatalf("failed to create log file: %s", err)
	} else {
		log.Out = file
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
	for {
		msgIn, err := stream.Recv()
		if err != nil {
			return err
		}
		if msgIn != nil {
			log.Debugf("RECV: %s", msgIn)
		}

		msgOut := &pb.ConnectResponse{}
		err = stream.Send(msgOut)
		if err != nil {
			log.Errorf("failed to send: %s", err)
		} else {
			log.Debugf("SEND: %s", msgOut)
		}
	}
}
