package app

import (
	"context"
	"github.com/x3a-tech/logit-go"
	"vacancy-parser/internal/config"
	"vacancy-parser/internal/sources"
)

type App struct {
	cfg    *config.Config
	src    *sources.Sources
	logger logit.Logger
}

type Params struct {
	Cfg    *config.Config
	Src    *sources.Sources
	Logger logit.Logger
}

func NewApp(params *Params) *App {
	return &App{
		cfg:    params.Cfg,
		src:    params.Src,
		logger: params.Logger,
	}
}

func (a *App) Init(ctx context.Context) error {
	const op = "app.Init"
	ctx = a.logger.NewOpCtx(ctx, op)

	err := a.src.Init(ctx)
	if err != nil {
		return err
	}

	return nil
}
