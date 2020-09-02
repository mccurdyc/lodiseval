package replicamanager

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type execFn func(context.Context, []string) error

func start(addr string, logger *log.Logger) execFn {
	return func(_ context.Context, _ []string) error {
		lis, err := net.Listen("tcp", addr)
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

		// Register ReplicaManager server.
		RegisterReplicaManagerSvcServer(s, &server{
			logger:   logger,
			replicas: make(map[string]string),
		})

		logger.Printf("replicamanager server listening at %s\n", addr)
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func createReplica(addr string, rmc ReplicaManagerSvcClient, alg string, store string, _ *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		_, err := rmc.CreateReplica(ctx, &CreateReplicaRequest{
			Address:         addr,
			Algorithm:       alg,
			Store:           store,
			SyncIntervalSec: 5,
		})

		return err
	}
}

func listReplicas(rmc ReplicaManagerSvcClient, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		lr, err := rmc.ListReplicas(ctx, &ListReplicasRequest{})

		fmt.Printf("REPLICAS: %d\n", len(lr.Replicas))
		for i, r := range lr.Replicas {
			fmt.Printf("%d\t%s\t%s\n", i, r.Id, r.Address)
		}

		return err
	}
}
