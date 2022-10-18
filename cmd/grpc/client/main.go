// command client
package main

import (
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

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer conn.Close()

	message := &pb.Message{}
	log.Infof("message: %v", message)

	log.Info("client: stopped")
}
