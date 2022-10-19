// command server
package main

import (
	"net"
	"os"

	"github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
}

func main() {
	log.Info("server: started")

	conn, err := net.Listen("tcp", "localhost:65000")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	server := grpc.NewServer()
	pb.RegisterConnectionServer(server, &connectionServer{})
	server.Serve(conn)

	log.Info("server: stopped")
}

type connectionServer struct {
	pb.UnimplementedConnectionServer
}

func (s *connectionServer) Connection(stream pb.Connection_ConnectServer) error {
	for {
		msgIn, err := stream.Recv()
		if err != nil {
			return err
		}
		if msgIn != nil {
			log.Infof("RECV MESSAGE: %s", msgIn)
		}

		body, err := anypb.New(&pb.Response{})
		msg := &pb.Message{
			RequestResponse: &pb.Message_Response{
				Response: body,
			},
		}
		stream.Send(msg)
		log.Infof("SEND MESSAGE: %s", msg)
	}
}
