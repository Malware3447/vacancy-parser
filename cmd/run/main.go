package run

import (
	"context"
	"github.com/x3a-tech/configo"
	"github.com/x3a-tech/envo"
	"github.com/x3a-tech/logit-go"
	"vacancy-parser/internal/config"
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
}
