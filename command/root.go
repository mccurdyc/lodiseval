package command

import (
	"context"
	"os"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func Run(ctx context.Context, args []string) int {
	root := &ffcli.Command{
		Name:        "lodiseval",
		ShortUsage:  "lodiseval [flags]",
		Subcommands: []*ffcli.Command{alg},
	}

	err := root.ParseAndRun(context.Background(), os.Args[1:])
	if err != nil {
		return 1
	}

	return 0
}
