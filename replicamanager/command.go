package replicamanager

import (
	"context"
	"flag"
	"log"

	"github.com/peterbourgon/ff/v3/ffcli"
	"google.golang.org/grpc"
)

var (
	mgrFlagSet = flag.NewFlagSet("replicamanager", flag.ExitOnError)
	addr       = mgrFlagSet.String("addr", "127.0.0.1:8118", "The replicamanager address.")

	replicaFlagSet = flag.NewFlagSet("replica", flag.ExitOnError)
	replicaAddr    = replicaFlagSet.String("addr", "127.0.0.1:8119", "The replica address.")
	alg            = replicaFlagSet.String("algorithm", "simpleleader", "The algorithm used by the replica.") // should this be on the manager instead?
	store          = replicaFlagSet.String("addr", "mapstore", "The backend store used by the replica.")

	serviceName = "ReplicaManager"
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand(logger *log.Logger) *ffcli.Command {
	// TODO: Make secure communication possible via TLS.
	// TODO: This connection currently doesn't get closed.
	conn, _ := grpc.Dial(*addr, grpc.WithInsecure())

	rmc := NewReplicaManagerSvcClient(conn)

	return &ffcli.Command{
		Name:       "replicamanager",
		ShortUsage: "lodiseval replicamanager <subcommand>",
		ShortHelp:  "Manage replicas in the cluster",
		FlagSet:    mgrFlagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "start",
				ShortUsage: "start [flags]",
				ShortHelp:  "Start the replicamanager server.",
				FlagSet:    mgrFlagSet,
				Exec:       start(*addr, logger),
			},
			{
				Name:       "replica",
				ShortUsage: "lodiseval replicamanager replica <subcommand>",
				ShortHelp:  "Interact with an individual replica in the cluster.",
				FlagSet:    mgrFlagSet,
				Subcommands: []*ffcli.Command{
					{
						Name:       "start",
						ShortUsage: "start [flags]",
						ShortHelp:  "Start a single replica server.",
						FlagSet:    replicaFlagSet,
						Exec:       createReplica(*replicaAddr, rmc, *alg, *store, logger),
					},
					{
						Name:       "list",
						ShortUsage: "list [flags]",
						ShortHelp:  "List all running replicas managed by the replicamanager.",
						FlagSet:    replicaFlagSet,
						Exec:       listReplicas(rmc, logger),
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
