package replicamanager

import (
	"context"
	"flag"
	"log"

	"github.com/peterbourgon/ff/v3/ffcli"
)

const DefaultManagerAddress = "127.0.0.1:8118"

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		mgrFlagSet = flag.NewFlagSet("replicamanager", flag.ExitOnError)
		addr       = mgrFlagSet.String("addr", DefaultManagerAddress, "The replicamanager address.")

		replicaFlagSet = flag.NewFlagSet("replicamanager replica", flag.ExitOnError)
	)

	return &ffcli.Command{
		Name:       "replicamanager",
		ShortUsage: "lodiseval replicamanager <subcommand>",
		ShortHelp:  "Manage replicas in the cluster",
		FlagSet:    mgrFlagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "start",
				ShortUsage: "lodiseval replicamanager start [flags]",
				ShortHelp:  "Start the replicamanager server.",
				FlagSet:    mgrFlagSet,
				Exec:       start(addr, logger),
			},
			{
				Name:       "replica",
				ShortUsage: "lodiseval replicamanager replica <subcommand>",
				ShortHelp:  "Gather information about replicas managed by the manager.",
				FlagSet:    replicaFlagSet,
				Exec: func(context.Context, []string) error {
					return flag.ErrHelp
				},
				Subcommands: []*ffcli.Command{
					{
						Name:       "list",
						ShortUsage: "lodiseval replicamanager replica list [flags]",
						ShortHelp:  "List all running replicas managed by the replicamanager.",
						FlagSet:    replicaFlagSet,
						Exec:       listReplicas(addr, logger),
					},
				},
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
