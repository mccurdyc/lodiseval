package replicamanager

import (
	"context"
	"log"
	sync "sync"
)

type server struct {
	UnimplementedReplicaManagerSvcServer

	logger *log.Logger

	// TODO: should these have separate mutexes?
	m        sync.Mutex
	replicas map[string]string
}

// =====================================================
// Replica Management.
// =====================================================
func (s *server) RegisterReplica(_ context.Context, req *RegisterReplicaRequest) (*RegisterReplicaResponse, error) {
	s.m.Lock()
	s.replicas[req.Replica.Id] = req.Replica.Address
	s.m.Unlock()

	return &RegisterReplicaResponse{}, nil
}

func (s *server) ListReplicas(_ context.Context, req *ListReplicasRequest) (*ListReplicasResponse, error) {
	var replicas []*Replica

	s.m.Lock()
	for k, v := range s.replicas {
		replicas = append(replicas, &Replica{
			Id:      k,
			Address: v,
		})
	}
	s.m.Unlock()

	return &ListReplicasResponse{
		Replicas: replicas,
	}, nil
}
