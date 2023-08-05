// Package grpc handles the grpc connection
package grpc

import (
	"fmt"
	"net"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	log "github.com/sirupsen/logrus"
	pb "github.com/tyholling/golang/proto/grpc/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Metrics stores metrics for prometheus
type Metrics struct {
	cpu         prometheus.Gauge
	memory      prometheus.Gauge
	bytesIn     prometheus.Counter
	bytesOut    prometheus.Counter
	errorsIn    prometheus.Counter
	errorsOut   prometheus.Counter
	discardsIn  prometheus.Counter
	discardsOut prometheus.Counter
}

var metrics = Metrics{
	cpu: promauto.NewGauge(
		prometheus.GaugeOpts{Name: "cpu_utilization", Help: "cpu utilization"}),
	memory: promauto.NewGauge(
		prometheus.GaugeOpts{Name: "memory_utilization", Help: "memory utilization"}),
	bytesIn: promauto.NewCounter(
		prometheus.CounterOpts{Name: "bytes_in_total", Help: "bytes in total"}),
	bytesOut: promauto.NewCounter(
		prometheus.CounterOpts{Name: "bytes_out_total", Help: "bytes out total"}),
	errorsIn: promauto.NewCounter(
		prometheus.CounterOpts{Name: "errors_in_total", Help: "errors in total"}),
	errorsOut: promauto.NewCounter(
		prometheus.CounterOpts{Name: "errors_out_total", Help: "errors out total"}),
	discardsIn: promauto.NewCounter(
		prometheus.CounterOpts{Name: "discards_in_total", Help: "discards in total"}),
	discardsOut: promauto.NewCounter(
		prometheus.CounterOpts{Name: "discards_out_total", Help: "discards out total"}),
}

// Server represents the grpc server
type Server struct {
	conn   net.Listener
	server *grpc.Server
}

// Listen listens for new grpc connections
func (s *Server) Listen() error {
	conn, err := net.Listen("tcp", "0.0.0.0:65000") // #nosec G102
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	s.conn = conn

	return nil
}

// Start accepts incoming grpc connections
func (s *Server) Start() {
	s.server = grpc.NewServer()
	pb.RegisterConnectionServiceServer(s.server, &connectionServer{})
	err := s.server.Serve(s.conn)
	if err != nil {
		log.Errorf("failed to start server: %s", err)
		return
	}
}

type connectionServer struct {
	pb.UnimplementedConnectionServiceServer
}

// Connect runs the handlers to send and receive grpc messages
func (s *connectionServer) Connect(stream pb.ConnectionService_ConnectServer) error {
	wg := sync.WaitGroup{}
	msgChan := make(chan *pb.ConnectResponse)

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleSend(stream, msgChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		handleRecv(stream, msgChan)
	}()

	// subscribe to heartbeat
	request1, err := anypb.New(&pb.Subscribe{
		Type: pb.SubscriptionType_HEARTBEAT,
	})
	if err != nil {
		log.Fatal(err)
	}
	msgChan <- &pb.ConnectResponse{
		Request: request1,
	}

	// subscribe to metrics
	request2, err := anypb.New(&pb.Subscribe{
		Type: pb.SubscriptionType_METRICS,
	})
	if err != nil {
		log.Fatal(err)
	}
	msgChan <- &pb.ConnectResponse{
		Request: request2,
	}

	wg.Wait()
	return nil
}

func handleSend(stream pb.ConnectionService_ConnectServer, msgChan <-chan *pb.ConnectResponse) {
	for msg := range msgChan {
		err := stream.Send(msg)
		if err != nil {
			log.Errorf("failed to send: %s", err)
			return
		}
		log.Debugf("send: %s", msg)
	}
}

func handleRecv(stream pb.ConnectionService_ConnectServer, _ chan<- *pb.ConnectResponse) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Errorf("failed to receive: %s", err)
			return
		}

		if msg.Request != nil {
			request, err := anypb.UnmarshalNew(msg.Request, proto.UnmarshalOptions{})
			if err != nil {
				log.Error(err)
				continue
			}
			log.Debugf("received request: %s", request)
		} else if msg.Response != nil {
			response, err := anypb.UnmarshalNew(msg.Response, proto.UnmarshalOptions{})
			if err != nil {
				log.Error(err)
				continue
			}
			if v, ok := response.(*pb.Metrics); ok {
				metrics.cpu.Set(v.Cpu)
				metrics.memory.Set(v.Memory)

				metrics.bytesIn.Add(float64(v.BytesIn))
				metrics.bytesOut.Add(float64(v.BytesOut))
				metrics.errorsIn.Add(float64(v.ErrorsIn))
				metrics.errorsOut.Add(float64(v.ErrorsOut))
				metrics.discardsIn.Add(float64(v.DiscardsIn))
				metrics.discardsOut.Add(float64(v.DiscardsOut))
			}
			log.Debugf("received response: %s", response)
		}
	}
}
