package command

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/mccurdyc/lodiseval/algorithm"
	"github.com/mccurdyc/lodiseval/builtin/algorithm/simpleleader"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var algorithms = map[string]algorithm.Factory{
	"simpleleader": simpleleader.Factory,
}

var (
	algFlagSet = flag.NewFlagSet("algorithm", flag.ExitOnError)
	stateStore = algFlagSet.String("store", "mapstore", "the backend state store")
	verbose    = algFlagSet.Bool("v", false, "increase log verbosity")

	alg = &ffcli.Command{
		Name:        "algorithm",
		ShortUsage:  "lodiseval algorithm <subcommand>",
		ShortHelp:   "Analyze concensus algorithms",
		FlagSet:     algFlagSet,
		Subcommands: []*ffcli.Command{algEval, algList, algDescribe},
	}

	algEval = &ffcli.Command{
		Name:       "eval",
		ShortUsage: "eval [flags]",
		ShortHelp:  "Evaluate a concensus algorithm",
		FlagSet:    algFlagSet,
		Exec:       algorithmEval,
	}

	algList = &ffcli.Command{
		Name:       "list",
		ShortUsage: "list [flags]",
		ShortHelp:  "List supported concensus algorithms",
		FlagSet:    algFlagSet,
		Exec:       algorithmList,
	}

	algDescribe = &ffcli.Command{
		Name:       "describe",
		ShortUsage: "describe <algorithm>",
		ShortHelp:  "Describe concensus algorithm",
		FlagSet:    algFlagSet,
		Exec:       algorithmDescribe,
	}
)

func algorithmEval(_ context.Context, _ []string) error {
	return errors.New("not implemented")
}

func algorithmList(_ context.Context, _ []string) error {
	for k := range algorithms {
		fmt.Printf("\t%s\n", k)
	}

	return nil
}

func algorithmDescribe(_ context.Context, _ []string) error {
	return errors.New("not implemented")
}
