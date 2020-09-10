package replica

import (
	"context"
	"flag"
	"log"

	"github.com/mccurdyc/lodiseval/replicamanager"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func NewCommand(logger *log.Logger) *ffcli.Command {
	var (
		replicaFlagSet = flag.NewFlagSet("replica", flag.ExitOnError)
		addr           = replicaFlagSet.String("addr", "127.0.0.1:8119", "The replica address.")
		mgrAddr        = replicaFlagSet.String("manager-addr", replicamanager.DefaultManagerAddress, "The replica manager address.")
		alg            = replicaFlagSet.String("algorithm", "simpleleader", "The algorithm used by the replica.") // should this be on the manager instead?
		store          = replicaFlagSet.String("store", "mapstore", "The backend store used by the replica.")
	)

	return &ffcli.Command{
		Name:       "replica",
		ShortUsage: "lodiseval replica <subcommand>",
		ShortHelp:  "Interact with an individual replica in the cluster.",
		FlagSet:    replicaFlagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "start",
				ShortUsage: "start [flags]",
				ShortHelp:  "Start a single replica server.",
				FlagSet:    replicaFlagSet,
				Exec:       start(addr, mgrAddr, alg, store, logger),
			},
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
}
