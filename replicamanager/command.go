package replicamanager

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/peterbourgon/ff/v3/ffcli"
	"google.golang.org/grpc"
)

var (
	flagSet = flag.NewFlagSet("replicamanager", flag.ExitOnError)
	port    = flagSet.Int("port", 8118, "the port where replicamanager will listen")
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand() *ffcli.Command {
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
				Exec:       start(*port),
			},
		},
	}
}

func start(port int) func(context.Context, []string) error {
	return func(_ context.Context, _ []string) error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			return err
		}

		s := grpc.NewServer()
		RegisterHealthServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}
