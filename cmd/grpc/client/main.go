// command client
package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Out = os.Stdout
}

func main() {
	log.Info("client: started")
	defer log.Info("client: stopped")

	target := "localhost:65000"
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("failed to connect to channel: %s", err)
		return
	}
	defer conn.Close()

	client := pb.NewConnectionClient(conn)
	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Errorf("failed to connect to server: %s", err)
		return
	}

	err = stream.Send(&pb.Message{
		RequestResponse: &pb.Message_Request{},
	})
	if err != nil {
		log.Errorf("failed to send message: %s", err)
		return
	}

	for {
		msgIn, err := stream.Recv()
		if err != nil {
			log.Errorf("failed to read message: %s", err)
			continue
		}
		if msgIn != nil {
			log.Infof("RECV MESSAGE: %s", msgIn)
		}

		msg := &pb.Message{
			RequestResponse: &pb.Message_Response{},
		}
		err = stream.Send(msg)
		if err != nil {
			log.Errorf("failed to send message: %s", err)
		} else {
			log.Infof("SEND MESSAGE: %s", msg)
		}
	}
}
