package command

import (
	"context"
	"errors"
	"flag"
	"fmt"

	mapstore "github.com/mccurdyc/lodiseval/builtin/store/map"
	"github.com/mccurdyc/lodiseval/store"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var (
	stores = map[string]store.Factory{
		"mapstore": mapstore.Factory,
	}
)

var (
	sFlagSet = flag.NewFlagSet("store", flag.ExitOnError)

	s = &ffcli.Command{
		Name:        "store",
		ShortUsage:  "lodiseval store <subcommand>",
		ShortHelp:   "Details about supported state stores",
		FlagSet:     sFlagSet,
		Subcommands: []*ffcli.Command{sList, sDescribe},
	}

	sList = &ffcli.Command{
		Name:       "list",
		ShortUsage: "list [flags]",
		ShortHelp:  "List supported state stores",
		FlagSet:    sFlagSet,
		Exec:       storeList,
	}

	sDescribe = &ffcli.Command{
		Name:       "describe",
		ShortUsage: "describe <store>",
		ShortHelp:  "Describe state store",
		FlagSet:    sFlagSet,
		Exec:       storeDescribe,
	}
)

func storeList(_ context.Context, _ []string) error {
	for k := range stores {
		fmt.Printf("\t%s\n", k)
	}

	return nil
}

func storeDescribe(_ context.Context, _ []string) error {
	return errors.New("not implemented")
}
