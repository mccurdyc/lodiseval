package replicamanager

import "context"

type server struct {
}

func (s *server) Health(_ context.Context, _ *HealthRequest) (*HealthResponse, error) {
	return &HealthResponse{
		Value: "OK",
	}, nil
}
