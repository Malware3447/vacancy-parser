package hh

import (
	"context"
	"fmt"
	"github.com/x3a-tech/logit-go"
	"net/http"
	"vacancy-parser/internal/config"
)

type HhParser struct {
	cfg    *config.Config
	logger logit.Logger
}

type HhParams struct {
	Cfg    *config.Config
	Logger logit.Logger
}

func NewHhParser(ctx context.Context, params HhParams) *HhParser {
	return &HhParser{
		cfg:    params.Cfg,
		logger: params.Logger,
	}
}

func (h *HhParser) LoadAndCollect(ctx context.Context, path string) error {
	const op = "parser.hh.LoadAndCollect"
	ctx = h.logger.NewOpCtx(ctx, op)

	_, _ = h.LoadDoc(ctx, path)
	return nil
}

func (h *HhParser) LoadDoc(ctx context.Context, path string) (string, error) {
	const op = "parser.hh.LoadDoc"
	ctx = h.logger.NewOpCtx(ctx, op)

	h.logger.Info(ctx, fmt.Sprintf("Начало загрузки страницы: %v", path))

	resp, err := http.Get(path)
	if err != nil {
		return "", fmt.Errorf("ошибка при получении страницы: %v", err)
	}
	defer resp.Body.Close()

	return "", nil
}
