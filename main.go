package main

import (
	"context"
	"os"

	"github.com/mccurdyc/lodiseval/command"
)

func main() {
	os.Exit(command.Run(context.Background(), os.Args[1:]))
}
