package algorithm

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var (
	ErrEmptyAlg       = errors.New("algorithm must be specified")
	ErrUnsupportedAlg = errors.New("unsupported algorithm")
)

// TODO: think a bit more as there is some implementation detail bleed with ffcli.Command being returned.
func NewCommand(builtins map[string]Factory) *ffcli.Command {
	flagSet := flag.NewFlagSet("algorithm", flag.ExitOnError)
	// stateStore := flagSet.String("store", "mapstore", "the backend state store")
	// verbose := flagSet.Bool("v", false, "increase log verbosity")

	return &ffcli.Command{
		Name:       "algorithm",
		ShortUsage: "lodiseval algorithm <subcommand>",
		ShortHelp:  "Analyze concensus algorithms",
		FlagSet:    flagSet,
		Subcommands: []*ffcli.Command{
			{
				Name:       "eval",
				ShortUsage: "eval [flags]",
				ShortHelp:  "Evaluate a concensus algorithm",
				FlagSet:    flagSet,
				Exec:       evalCmd(builtins),
			},
			{
				Name:       "list",
				ShortUsage: "list [flags]",
				ShortHelp:  "List supported concensus algorithms",
				FlagSet:    flagSet,
				Exec:       listCmd(builtins),
			},
			{
				Name:       "describe",
				ShortUsage: "describe <algorithm>",
				ShortHelp:  "Describe concensus algorithm",
				FlagSet:    flagSet,
				Exec:       describeCmd(builtins),
			},
		},
	}
}

func evalCmd(builtins map[string]Factory) func(context.Context, []string) error {
	return func(ctx context.Context, args []string) error {
		_, err := parseAlgorithm(ctx, args, builtins)
		if err != nil {
			return err
		}

		// TODO: This may be where in a separate Goroutine, we "Start" the algorithm
		// and send it input via some input stream; internally using channels.
		//
		// This to me indicates a client-server architecture.
		// This is also a good opportunity just to mess with gRPC for fun.
		return nil
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

func describeCmd(builtins map[string]Factory) func(context.Context, []string) error {
	return func(ctx context.Context, args []string) error {
		alg, err := parseAlgorithm(ctx, args, builtins)
		if err != nil {
			return err
		}

		// TODO: I don't really want to write here, but instead would prefer centralizing output.
		fmt.Printf("%s\n\t%s\n", args[0], alg.Describe())

		return nil
	}
}

func parseAlgorithm(ctx context.Context, args []string, builtins map[string]Factory) (Algorithm, error) {
	if len(args) < 1 {
		return nil, ErrEmptyAlg
	}

	algFactory, ok := builtins[args[0]]
	if !ok {
		return nil, ErrUnsupportedAlg
	}

	// TODO: Create and parse flags for setting values in the config.
	cfg := Config{}

	alg, err := algFactory(ctx, &cfg)
	if err != nil {
		return nil, errors.New("error creating algorithm")
	}

	return alg, nil
}
