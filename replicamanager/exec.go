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

const serviceName = "ReplicaManager"

type execFn func(context.Context, []string) error

func start(addr *string, logger *log.Logger) execFn {
	return func(_ context.Context, _ []string) error {
		lis, err := net.Listen("tcp", *addr)
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

		logger.Printf("replicamanager server listening at %s\n", *addr)
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func listReplicas(mgrAddr *string, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
		// TODO: Make secure communication possible via TLS.
		// TODO: This connection currently doesn't get closed.
		conn, _ := grpc.Dial(*mgrAddr, grpc.WithInsecure())
		rmc := NewReplicaManagerSvcClient(conn)

		lr, err := rmc.ListReplicas(ctx, &ListReplicasRequest{})

		fmt.Printf("REPLICAS: %d\n", len(lr.Replicas))
		for i, r := range lr.Replicas {
			fmt.Printf("%d\t%s\t%s\n", i, r.Id, r.Address)
		}

		return err
	}
}
