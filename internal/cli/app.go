package cli

import (
	"github.com/urfave/cli/v3"

	"github.com/Lucashcr/changelogr/internal/cli/commands"
)

func New(version string) *cli.Command {
	return &cli.Command{
		Name:    "changelogr",
		Usage:   "A application changelog builder",
		Version: version,
		Commands: []*cli.Command{
			commands.NewInitCommand(),
		},
	}
}
