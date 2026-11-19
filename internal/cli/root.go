package cli

import (
	urfave "github.com/urfave/cli/v3"

	"github.com/Lucashcr/changelogr/internal/app"
	"github.com/Lucashcr/changelogr/internal/cli/commands"
)

func New(version string) *urfave.Command {
	env := &app.Env{}

	return &urfave.Command{
		Name:    "changelogr",
		Usage:   "A application changelog builder",
		Version: version,

		Before: env.Load,

		EnableShellCompletion: true,
		ConfigureShellCompletionCommand: func(cmd *urfave.Command) {
			cmd.Hidden = false
			cmd.Usage = "Generate shell completion scripts"
		},

		Commands: []*urfave.Command{
			commands.NewInit(env),
		},
	}
}
