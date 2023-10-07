package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", "localhost:65000")
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.listener = listener
	s.server = grpc.NewServer()

	err = s.server.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

func main() {
	server := &Server{}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
