package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/rs/zerolog"

	"github.com/mccurdyc/lodiseval/algorithm"
	"github.com/mccurdyc/lodiseval/builtin/algorithm/simpleleader"
	"github.com/mccurdyc/lodiseval/builtin/store/mapstore"
	"github.com/mccurdyc/lodiseval/replica"
	"github.com/mccurdyc/lodiseval/replicamanager"
	"github.com/mccurdyc/lodiseval/store"
)

func main() {
	l := zerolog.New(os.Stdout)
	logger := log.New(l, "", log.Ldate|log.Ltime|log.LUTC)

	os.Exit(Run(context.Background(), os.Args[1:], logger))
}

func Run(ctx context.Context, args []string, l *log.Logger) int {
	var (
		rootFlagSet = flag.NewFlagSet("lodiseval", flag.ExitOnError)
	)

	root := &ffcli.Command{
		Name:       "lodiseval",
		ShortUsage: "lodiseval <subcommand> [flags]",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			replicamanager.NewCommand(l),
			replica.NewCommand(l),
			algorithm.NewCommand(
				map[string]algorithm.Factory{
					"simpleleader": simpleleader.Factory,
				},
			),
			store.NewCommand(
				map[string]store.Factory{
					"mapstore": mapstore.Factory,
				},
			),
		},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		l.Fatal(err)
		return 1
	}

	return 0
}
