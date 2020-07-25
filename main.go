package main

import (
	"context"
	"log"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"

	"github.com/mccurdyc/lodiseval/algorithm"
	"github.com/mccurdyc/lodiseval/builtin/algorithm/simpleleader"
	"github.com/mccurdyc/lodiseval/builtin/store/mapstore"
	"github.com/mccurdyc/lodiseval/replicamanager"
	"github.com/mccurdyc/lodiseval/store"
)

func main() {
	os.Exit(Run(context.Background(), os.Args[1:]))
}

func Run(ctx context.Context, args []string) int {
	root := &ffcli.Command{
		Name:       "lodiseval",
		ShortUsage: "lodiseval <subcommand> [flags]",
		Subcommands: []*ffcli.Command{
			replicamanager.NewCommand(),
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
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		// TODO better logging
		log.Fatal(err)
		return 1
	}

	return 0
}
