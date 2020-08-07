package replica

import (
	"context"
	"log"
)

type server struct {
	UnimplementedReplicaServer

	logger *log.Logger

	id string
}

func (s *server) ID(ctx context.Context, req *IDRequest) (*IDResponse, error) {
	return &IDResponse{
		Id: s.id,
	}, nil
}
