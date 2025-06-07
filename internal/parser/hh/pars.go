package hh

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/x3a-tech/logit-go"
	"net/http"
	"vacancy-parser/internal/config"
	"vacancy-parser/internal/models"
)

type HhParser struct {
	cfg    *config.Config
	query  *models.ListParamsQuery
	logger logit.Logger
}

type HhParams struct {
	Cfg    *config.Config
	Query  *models.ListParamsQuery
	Logger logit.Logger
}

func NewHhParser(ctx context.Context, params HhParams) *HhParser {
	return &HhParser{
		cfg:    params.Cfg,
		query:  params.Query,
		logger: params.Logger,
	}
}

func (h *HhParser) LoadAndCollect(ctx context.Context, path string) error {
	const op = "parser.hh.LoadAndCollect"
	ctx = h.logger.NewOpCtx(ctx, op)

	doc, err := h.LoadDoc(ctx, path)
	if err != nil {
		return fmt.Errorf("ошибка при загрузке документа: %v", err)
	}

	_, err = h.getItems(ctx, doc)
	if err != nil {
		return fmt.Errorf("ошибка при сборе данных: %v", err)
	}
	return nil
}

func (h *HhParser) LoadDoc(ctx context.Context, path string) (*goquery.Document, error) {
	const op = "parser.hh.LoadDoc"
	ctx = h.logger.NewOpCtx(ctx, op)

	h.logger.Info(ctx, fmt.Sprintf("Начало загрузки страницы: %v", path))

	resp, err := http.Get(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении страницы: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			h.logger.Error(ctx, fmt.Errorf("ошибка при закрытии тела ответа: %v", err))
		}
	}()

	if resp.StatusCode != 200 {
		h.logger.Error(ctx, fmt.Errorf("неожиданный статус ответа: %v", resp.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	h.logger.Info(ctx, fmt.Sprintf("Загрузка страницы завершена: %v", path))

	return doc, nil
}

func (h *HhParser) getItems(ctx context.Context, doc *goquery.Document) ([]*models.ItemsModel, error) {
	const op = "parser.hh.getItems"
	ctx = h.logger.NewOpCtx(ctx, op)

	selection := models.ItemsList{}
	items := make([]*models.ItemsModel, 0)

	selection.Items = doc.Find(h.query.Items)

	errorMap := make(map[int]error)
	selection.Items.Each(func(index int, selection *goquery.Selection) {

		item := ParseListItems(
			index,
			selection,
			h.query,
			errorMap,
		)

		if item != nil {
			items = append(items, item)
		}

	})

	if len(errorMap) > 0 {
		err := fmt.Errorf("ошибки при парсинге: %+v", errorMap)
		return nil, err
	}

	return items, nil
}
