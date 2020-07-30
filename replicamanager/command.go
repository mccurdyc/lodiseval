package replicamanager

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/peterbourgon/ff/v3/ffcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	flagSet = flag.NewFlagSet("replicamanager", flag.ExitOnError)
	port    = flagSet.Int("port", 8118, "the port where replicamanager will listen")

	serviceName = "ReplicaManager"
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand(logger *log.Logger) *ffcli.Command {
	return &ffcli.Command{
		Name:       "replicamanager",
		ShortUsage: "lodiseval replicamanager <subcommand>",
		ShortHelp:  "Manage replicas in the cluster",
		FlagSet:    flagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "start",
				ShortUsage: "start [flags]",
				ShortHelp:  "Start the replicamanager",
				FlagSet:    flagSet,
				Exec:       start(*port, logger),
			},
		},
	}
}

func start(port int, logger *log.Logger) func(context.Context, []string) error {
	return func(_ context.Context, _ []string) error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return err
		}

		s := grpc.NewServer()

		// Add standardized healthcheck.
		healthcheck := health.NewServer()
		healthpb.RegisterHealthServer(s, healthcheck)

		// Add reflection so that clients can query for available services, methods, etc.
		reflection.Register(s)

		// Register ReplicaManager server.
		RegisterReplicaManagerServer(s, &server{})

		// https://github.com/grpc/grpc-go/blob/master/examples/features/health/server/main.go
		healthcheck.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)

		logger.Printf("server listening on port :%d\n", port)
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}
