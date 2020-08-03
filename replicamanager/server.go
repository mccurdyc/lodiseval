package replicamanager

import (
	"context"
	"log"

	"github.com/mccurdyc/lodiseval/replica"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type server struct {
	UnimplementedReplicaManagerServer

	logger *log.Logger
}

// =====================================================
// Replica Management.
// =====================================================
func (s *server) CreateReplica(ctx context.Context, req *CreateReplicaRequest) (*CreateReplicaResponse, error) {
	cfg := replica.Config{
		Addr:   req.Address,
		Logger: s.logger,
	}

	err := replica.Create(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	return nil, status.Errorf(codes.Unimplemented, "method CreateReplica not implemented")
}

// func (s *server) ListReplicas(_ context.Context, _ *ListReplicasRequest) *ListReplicasResponse {
// }
//
// func (s *server) DeleteReplica(_ context.Context, _ *DeleteReplicaRequest) *DeleteReplicaResponse {
// }
//
// // =====================================================
// // Leader Management.
// // =====================================================
//
// func (s *server) SetLeader(_ context.Context, _ *SetLeaderRequest) *SetLeaderResponse {
// }
//
// func (s *server) GetLeader(_ context.Context, _ *GetLeaderRequest) *GetLeaderResponse {
// }
//
// // =====================================================
// // Value Storage and Retrieval Operations.
// // =====================================================
//
// func (s *server) Set(_ context.Context, _ *SetRequest) *SetResponse {
// }
//
// func (s *server) Get(_ context.Context, _ *GetRequest) *GetResponse {
// }
//
// // =====================================================
// // Algorithm Management.
// // =====================================================
//
// func (s *server) SetAlgorithm(_ context.Context, _ *SetAlgorithmRequest) *SetAlgorithmResponse {
// }
//
// func (s *server) GetAlgorithm(_ context.Context, _ *GetAlgorithmRequest) *GetAlgorithmResponse {
// }
