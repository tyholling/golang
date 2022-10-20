// command client
package main

import (
	"context"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}

	err := os.MkdirAll("log", 0o0644)
	if err != nil {
		log.Fatalf("failed to create log directory: %s", err)
	}
	file, err := os.Create("log/client.log")
	if err != nil {
		log.Fatalf("failed to create log file: %s", err)
	} else {
		log.Out = file
	}
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

	wg := sync.WaitGroup{}
	messageChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range messageChan {
			msg := &pb.ConnectRequest{}
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

	messageChan <- struct{}{}
	wg.Wait()
}
