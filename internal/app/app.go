package app

import (
	"context"
	"github.com/x3a-tech/logit-go"
	"vacancy-parser/internal/config"
)

type App struct {
	cfg    *config.Config
	logger logit.Logger
}

type Params struct {
	cfg    *config.Config
	logger logit.Logger
}

func NewApp(params *Params) *App {
	return &App{
		cfg:    params.cfg,
		logger: params.logger,
	}
}

func (a *App) Init(ctx context.Context) error {
	const op = "app.Init"
	ctx = a.logger.NewOpCtx(ctx, op)

	return nil
}
