// command client
package main

import (
	"context"
	"os"
	"time"

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

	client := pb.NewConnectionServiceClient(conn)
	stream, err := client.Connect(context.Background())
	if err != nil {
		log.Errorf("failed to connect to server: %s", err)
		return
	}

	err = stream.Send(&pb.ConnectRequest{})
	if err != nil {
		log.Errorf("failed to send: %s", err)
		return
	}

	delay := time.Millisecond
	for {
		msg := &pb.ConnectResponse{}
		if stream != nil {
			msg, err = stream.Recv()
		}
		if stream == nil || err != nil {
			log.Errorf("failed to receive: %s", err)

			log.Infof("reconnecting after delay: %v", delay)
			time.Sleep(delay)

			if delay < time.Minute {
				// increase backoff
				delay *= 2
				if delay > time.Minute {
					delay = time.Minute
				}
			}

			err = conn.Close()
			if err != nil {
				log.Warningf("failed to close connection: %s", err)
			}

			conn, err = grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Errorf("failed to connect to channel: %s", err)
				continue
			}
			log.Infof("connected to channel: %v", conn)

			client = pb.NewConnectionServiceClient(conn)
			stream, err = client.Connect(context.Background())
			if err != nil {
				log.Errorf("failed to connect to server: %s", err)
				continue
			}
			log.Infof("connected to server: %v", stream)

			err = stream.Send(&pb.ConnectRequest{})
			if err != nil {
				log.Errorf("failed to send: %s", err)
			}

			continue
		}
		delay = time.Millisecond // reset backoff

		if msg != nil {
			log.Infof("RECV: %s", msg)
		}

		msgOut := &pb.ConnectRequest{}
		err = stream.Send(msgOut)
		if err != nil {
			log.Errorf("failed to send: %s", err)
		} else {
			log.Infof("SEND: %s", msg)
		}
	}
}
