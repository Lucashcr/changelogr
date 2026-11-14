package app

import (
	"context"
	"os"
	"path/filepath"

	urfave "github.com/urfave/cli/v3"

	"github.com/Lucashcr/changelogr/internal/config"
)

type Env struct {
	cwd    string
	config config.Config
}

func (env Env) CWD() string {
	return env.cwd
}

func (env Env) Config() config.Config {
	return env.config
}

func (env Env) ConfigFileExists() (bool, error) {
	entries, err := filepath.Glob(filepath.Join(env.cwd, "changelogr.*"))
	if err != nil {
		return false, err
	}
	return len(entries) > 0, nil
}

func (env *Env) Load(ctx context.Context, _ *urfave.Command) (context.Context, error) {
	var err error

	env.cwd, err = os.Getwd()
	if err != nil {
		return ctx, err
	}

	env.config, err = config.LoadConfig(env.cwd)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
