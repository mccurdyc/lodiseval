package replica

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var serviceName = "replica"

type Config struct {
	ID     string
	Addr   string
	Logger *log.Logger
}

func Create(ctx context.Context, cfg *Config) error {
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	// Add standardized healthcheck.
	healthcheck := health.NewServer()
	healthcheck.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(s, healthcheck)

	// Add reflection so that clients can query for available services, methods, etc.
	reflection.Register(s)

	RegisterReplicaSvcServer(s, &server{
		id: cfg.ID,
	})

	cfg.Logger.Printf("replica server listening on %s\n", cfg.Addr)
	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
