// command server
package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc"
	"google.golang.org/grpc"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
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
			log.Infof("RECV: %s", msgIn)
		}

		msgOut := &pb.ConnectResponse{}
		err = stream.Send(msgOut)
		if err != nil {
			log.Errorf("failed to send: %s", err)
		} else {
			log.Infof("SEND: %s", msgOut)
		}
	}
}
