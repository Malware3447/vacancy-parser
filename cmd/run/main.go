package main

import (
	"context"
	"fmt"
	"github.com/x3a-tech/configo"
	"github.com/x3a-tech/envo"
	"github.com/x3a-tech/logit-go"
	"github.com/x3a-tech/spg"
	"vacancy-parser/internal/app"
	"vacancy-parser/internal/config"
	repoPG "vacancy-parser/internal/db/repository/pg"
	"vacancy-parser/internal/parser"
	"vacancy-parser/internal/services/db/pg"
	"vacancy-parser/internal/sources"
)

func main() {
	const op = "cmd.run.main"

	cfg, _ := configo.MustLoad[config.Config]()
	env, err := envo.New(cfg.Env())
	logParams := logit.Params{
		AppConf:    &cfg.App,
		LoggerConf: &cfg.Logger,
		SenConf:    &cfg.Sentry,
		Env:        (*configo.Env)(env),
	}

	logger := logit.MustNewLogger(&logParams)

	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	ctx := logger.NewCtx(context.Background(), op, nil)

	logger.Info(ctx, "Сервис запущен успешно")

	poolPg, err := spg.NewClient(ctx, &cfg.DatabasePg)
	if err != nil {
		logger.Fatal(ctx, fmt.Errorf("ошибка при запуске Postgres: %s", err))
		panic(err)
	}
	logger.Info(ctx, "Postgres успешно запущен")

	repoPg := repoPG.NewRepositoryPg(poolPg)

	pgService := pg.NewService(repoPg)

	ParserParams := parser.Params{
		Cfg:    cfg,
		Repo:   pgService,
		Logger: logger,
	}

	prs := parser.NewParser(ctx, &ParserParams)

	sourceParams := sources.Params{
		Prs:    prs,
		Logger: logger,
	}

	src := sources.NewSources(&sourceParams)

	appParams := app.Params{
		Cfg:    cfg,
		Src:    src,
		Logger: logger,
	}

	ap := app.NewApp(&appParams)

	err = ap.Init(ctx)
	if err != nil {
		logger.Fatal(ctx, fmt.Errorf("ошибка при инициализации приложения: %s", err))
		panic(err)
	}

}
