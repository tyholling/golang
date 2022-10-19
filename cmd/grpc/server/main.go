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
	log.Level = logrus.FatalLevel
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
	pb.RegisterConnectionServer(server, &connectionServer{})
	err = server.Serve(conn)
	if err != nil {
		log.Errorf("failed to start server: %s", err)
		return
	}
}

type connectionServer struct {
	pb.UnimplementedConnectionServer
}

func (s *connectionServer) Connect(stream pb.Connection_ConnectServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}
		if msg != nil {
			log.Infof("RECV MESSAGE: %s", msg)
		}

		msg.RequestResponse = &pb.Message_Response{}
		err = stream.Send(msg)
		if err != nil {
			log.Errorf("failed to send message: %s", err)
		} else {
			log.Infof("SEND MESSAGE: %s", msg)
		}
	}
}
