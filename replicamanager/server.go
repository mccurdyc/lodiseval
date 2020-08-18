package replicamanager

import (
	"context"
	"log"
	"strconv"
	sync "sync"

	"github.com/mccurdyc/lodiseval/replica"
)

type server struct {
	UnimplementedReplicaManagerSvcServer

	logger *log.Logger

	// TODO: should these have separate mutexes?
	m        sync.Mutex
	idCount  int
	replicas map[string]string
}

// =====================================================
// Replica Management.
// =====================================================
func (s *server) CreateReplica(ctx context.Context, req *CreateReplicaRequest) (*CreateReplicaResponse, error) {
	id := -1
	s.m.Lock()
	s.idCount += 1
	id = s.idCount
	s.m.Unlock()
	idStr := strconv.Itoa(id)

	cfg := replica.Config{
		ID:     idStr,
		Addr:   req.Address,
		Logger: s.logger,
	}

	err := replica.Create(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	s.m.Lock()
	s.replicas[idStr] = cfg.Addr
	s.m.Unlock()

	return &CreateReplicaResponse{
		Id: idStr,
	}, nil
}

// TODO NEXT!
// func (s *server) ListReplicas(_ context.Context, _ *ListReplicasRequest) (*ListReplicasResponse, error) {
// 	s.m.Lock()
// 	for id, addr := range s.replicas {
// 	}
// 	s.m.Unlock()
//
// 	return &ListReplicasResponse{
// 		Id: idStr,
// 	}, nil
// }
//
// func (s *server) printFormattedReplicaList() string {
//
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
