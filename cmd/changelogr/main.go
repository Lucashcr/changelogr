package main

import (
	"context"
	"os"

	"github.com/Lucashcr/changelogr/internal/cli"
)

var version = "dev"

func main() {
	cli.New(version).Run(context.Background(), os.Args)
}
