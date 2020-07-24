package replicamanager

import (
	"context"
	"errors"
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand() *ffcli.Command {
	flagSet := flag.NewFlagSet("replicamanager", flag.ExitOnError)

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
				Exec:       start(),
			},
		},
	}
}

func start() func(context.Context, []string) error {
	return func(_ context.Context, _ []string) error {
		return errors.New("not implemented")
	}
}
