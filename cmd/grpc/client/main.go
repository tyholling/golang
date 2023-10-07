package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	connection *grpc.ClientConn
}

func (c *Client) Connect() error {
	target := "localhost:65000"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	connection, err := grpc.Dial(target, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	c.connection = connection
	log.Print("client connected")

	return nil
}

func main() {
	client := &Client{}
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
}
