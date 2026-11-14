package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Lucashcr/changelogr/internal/cli"
)

var version = "dev"

func main() {
	if err := cli.Build(version).Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
