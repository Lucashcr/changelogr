package commands

import (
	"context"
	"fmt"
	"slices"
	"strings"

	urfave "github.com/urfave/cli/v3"

	"github.com/Lucashcr/changelogr/internal/app"
	"github.com/Lucashcr/changelogr/internal/config"
)

func handleInit(env *app.Env, cmd *urfave.Command) error {
	configFileExists, err := env.ConfigFileExists()
	if err != nil {
		return urfave.Exit(
			fmt.Sprintf("Failed to check config file: %s", err),
			EXIT_GENERAL_ERROR,
		)
	}

	if configFileExists {
		return urfave.Exit(
			"Cannot initialize: Config file already exists",
			EXIT_ALREADY_EXISTS,
		)
	}

	configType := cmd.String("type")
	if !slices.Contains(config.ConfigTypeOptions, configType) {
		return urfave.Exit(
			fmt.Sprintf(
				"Invalid type: %s. Should be one of: %s",
				configType,
				strings.Join(config.ConfigTypeOptions, ","),
			),
			EXIT_INVALID_INPUT,
		)
	}

	err = config.NewConfig().Save(env.CWD(), configType)
	if err != nil {
		return urfave.Exit(
			fmt.Sprintf("Failed to initialize: %s", err),
			EXIT_GENERAL_ERROR,
		)
	}

	return nil
}

func NewInitCommand(env *app.Env) *urfave.Command {
	return &urfave.Command{
		Name:  "init",
		Usage: "Initialize a new configuration",

		Action: func(_ context.Context, cmd *urfave.Command) error {
			return handleInit(env, cmd)
		},

		Flags: []urfave.Flag{
			&urfave.StringFlag{
				Name:     "type",
				Aliases:  []string{"t"},
				Required: false,
				OnlyOnce: true,
				Value:    "yaml",
			},
		},
	}
}
