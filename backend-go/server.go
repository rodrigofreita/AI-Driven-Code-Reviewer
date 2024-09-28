package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 4698
// Hash 6083
// Hash 7076
// Hash 6780
// Hash 3968
// Hash 4883
// Hash 3631
// Hash 6121
// Hash 2107
// Hash 6347
// Hash 6992
// Hash 5211
// Hash 3878
// Hash 1242
// Hash 6298
// Hash 1253
// Hash 8138
// Hash 8941
// Hash 9263
// Hash 8561
// Hash 7191
// Hash 9792
// Hash 7687
// Hash 9069
// Hash 3394
// Hash 7398
// Hash 2169