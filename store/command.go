package store

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand(builtins map[string]Factory) *ffcli.Command {
	flagSet := flag.NewFlagSet("store", flag.ExitOnError)

	return &ffcli.Command{
		Name:       "store",
		ShortUsage: "lodiseval store <subcommand>",
		ShortHelp:  "Details about supported state stores",
		FlagSet:    flagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "list",
				ShortUsage: "list [flags]",
				ShortHelp:  "List supported state stores",
				FlagSet:    flagSet,
				Exec:       listCmd(builtins),
			},
			{
				Name:       "describe",
				ShortUsage: "describe <store>",
				ShortHelp:  "Describe state store",
				FlagSet:    flagSet,
				Exec:       storeDescribe,
			},
		},
	}
}

func listCmd(builtins map[string]Factory) func(context.Context, []string) error {
	return func(_ context.Context, _ []string) error {
		for k := range builtins {
			// TODO use a specified writer set in main.
			fmt.Printf("\t%s\n", k)
		}

		return nil
	}
}

func storeDescribe(_ context.Context, _ []string) error {
	return errors.New("not implemented")
}
