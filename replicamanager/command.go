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
				ShortHelp:  "Start the replicamanager server.",
				FlagSet:    flagSet,
				Exec:       start(*port, logger),
			},
			{
				Name:       "replica",
				ShortUsage: "lodiseval replicamanager replica <subcommand>",
				ShortHelp:  "Interact with an individual replica in the cluster.",
				FlagSet:    flagSet,
				Subcommands: []*ffcli.Command{
					{
						Name:       "start",
						ShortUsage: "start [flags]",
						ShortHelp:  "Start a single replica server.",
						FlagSet:    flagSet,
						Exec: func(ctx context.Context, args []string) error {
							// TODO: Make secure communication possible.
							conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
							if err != nil {
								return err
							}
							defer conn.Close()

							rm := NewReplicaManagerClient(conn)
							_, err = rm.CreateReplica(ctx, &CreateReplicaRequest{
								// TODO: don't hardcode these values.
								Address:         ":8119",
								Algorithm:       "simpleleader",
								Store:           "mapstore",
								SyncIntervalSec: 5,
							})

							return err
						},
					},
					{
						Name:       "list",
						ShortUsage: "list [flags]",
						ShortHelp:  "List all running replicas managed by the replicamanager.",
						FlagSet:    flagSet,
						Exec: func(ctx context.Context, args []string) error {
							// TODO: Make secure communication possible.
							conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithInsecure())
							if err != nil {
								return err
							}
							defer conn.Close()

							rm := NewReplicaManagerClient(conn)
							_, err = rm.ListReplicas(ctx, &ListReplicasRequest{})
							return err
						},
					},
				},
				Exec: func(context.Context, []string) error {
					return flag.ErrHelp
				},
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
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
		healthcheck.SetServingStatus(serviceName, healthpb.HealthCheckResponse_SERVING)
		healthpb.RegisterHealthServer(s, healthcheck)

		// Add reflection so that clients can query for available services, methods, etc.
		reflection.Register(s)

		// Register ReplicaManager server.
		RegisterReplicaManagerServer(s, &server{
			logger:   logger,
			replicas: make(map[string]string),
		})

		logger.Printf("replicamanager server listening on port :%d\n", port)
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}
