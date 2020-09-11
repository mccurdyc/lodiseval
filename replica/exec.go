package replica

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/mccurdyc/lodiseval/replicamanager"
)

const serviceName = "Replica"

type execFn func(context.Context, []string) error

func start(addr *string, mgrAddr *string, alg *string, store *string, logger *log.Logger) execFn {
	return func(ctx context.Context, _ []string) error {
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

		// TODO: Make secure communication possible via TLS.
		// TODO: This connection currently doesn't get closed.
		conn, _ := grpc.Dial(*mgrAddr, grpc.WithInsecure())
		rmc := replicamanager.NewReplicaManagerSvcClient(conn)

		// This registers the replica with the replicamanager before successfully
		// starting the replica server. It is possible that the replica server fails
		// to start. The expectation is that the replicamanager server will keep
		// track of alive replicas via healthchecks.
		_, err = rmc.RegisterReplica(ctx, &replicamanager.RegisterReplicaRequest{
			Replica: &replicamanager.Replica{
				Id:      generateID(*addr),
				Address: *addr,
			}})

		if err != nil {
			err = fmt.Errorf("failed to register replica with replicamanager %w", err)
			logger.Print(err)
			return err
		}

		RegisterReplicaSvcServer(s, &server{
			logger: logger,
		})

		logger.Printf("replica server listening at %s\n", *addr)
		logger.Printf("replica server pointed at replicamanager running at %s\n", *mgrAddr)
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func generateID(addr string) string {
	h := md5.New()
	h.Write([]byte(addr))
	return hex.EncodeToString(h.Sum(nil))
}
