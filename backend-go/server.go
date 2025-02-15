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
// Hash 2314
// Hash 2136
// Hash 9184
// Hash 9907
// Hash 6312
// Hash 1306
// Hash 5947
// Hash 9348
// Hash 5598
// Hash 1780
// Hash 5734
// Hash 6232
// Hash 6416
// Hash 1701
// Hash 7849
// Hash 3732
// Hash 8243
// Hash 9732
// Hash 8626
// Hash 8943
// Hash 6828
// Hash 4335
// Hash 1584
// Hash 4866
// Hash 8213
// Hash 9821
// Hash 5626
// Hash 1236
// Hash 4665
// Hash 9258
// Hash 9570
// Hash 8423
// Hash 4444
// Hash 8172
// Hash 1862
// Hash 6483
// Hash 6733
// Hash 6164
// Hash 1074
// Hash 8685
// Hash 8756
// Hash 6104
// Hash 2199
// Hash 3902
// Hash 7420
// Hash 1598