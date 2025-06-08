package sources

import (
	"context"
	"fmt"
	"github.com/x3a-tech/logit-go"
	"vacancy-parser/internal/parser"
)

type Sources struct {
	prs    parser.Parser
	logger logit.Logger
}

type Params struct {
	Prs    parser.Parser
	Logger logit.Logger
}

func NewSources(params Params) *Sources {
	return &Sources{
		prs:    params.Prs,
		logger: params.Logger,
	}
}

func (s *Sources) Init(ctx context.Context) error {
	const op = "sources.Init"
	ctx = s.logger.NewOpCtx(ctx, op)

	s.Parse(ctx)

	return nil
}

func (s *Sources) Parse(ctx context.Context) error {
	const op = "sources.Parse"
	ctx = s.logger.NewOpCtx(ctx, op)

	hhUrl := "https://voronezh.hh.ru/search/vacancy?hhtmFrom=main&hhtmFromLabel=vacancy_search_line&enable_snippets=false&L_save_area=true&experience=between1And3&search_field=name&search_field=company_name&search_field=description&text=golang"

	NewHh(hhUrl)

	err := s.prs.LoadAndCollect(ctx, hhUrl)
	if err != nil {
		err = fmt.Errorf("ошибка при парсинге hh.ru: %v", err)
		return err
	}

	return nil
}
