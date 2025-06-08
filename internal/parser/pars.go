package parser

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/x3a-tech/logit-go"
	"net/http"
	"vacancy-parser/internal/config"
	"vacancy-parser/internal/models"
	"vacancy-parser/internal/services/db/pg"
)

type Parser struct {
	cfg    *config.Config
	repo   *pg.Service
	logger logit.Logger
}

type Params struct {
	Cfg    *config.Config
	Repo   *pg.Service
	Logger logit.Logger
}

func NewParser(ctx context.Context, params *Params) *Parser {
	return &Parser{
		cfg:    params.Cfg,
		repo:   params.Repo,
		logger: params.Logger,
	}
}

func (h *Parser) LoadAndCollect(ctx context.Context, itemParams *models.ItemParams) error {
	const op = "parser.hh.LoadAndCollect"
	ctx = h.logger.NewOpCtx(ctx, op)

	doc, err := h.LoadDoc(ctx, itemParams.BaseUrl)
	if err != nil {
		return fmt.Errorf("ошибка при загрузке документа: %v", err)
	}

	items, err := h.getItems(ctx, doc, itemParams)
	if err != nil {
		return fmt.Errorf("ошибка при сборе данных: %v", err)
	}

	err = h.UpdateVacancies(ctx, items)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении вакансий в БД: %v", err)
	}

	return nil
}

func (h *Parser) LoadDoc(ctx context.Context, path string) (*goquery.Document, error) {
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

	h.logger.Info(ctx, fmt.Sprintf("Загрузка страницы завершена"))

	return doc, nil
}

func (h *Parser) getItems(ctx context.Context, doc *goquery.Document, itemParams *models.ItemParams) ([]*models.ItemsModel, error) {
	const op = "parser.hh.getItems"
	ctx = h.logger.NewOpCtx(ctx, op)

	selections := models.ItemsList{}
	items := make([]*models.ItemsModel, 0)

	selections.Items = doc.Find(itemParams.Query.Items)

	errorMap := make(map[int]error)
	selections.Items.Each(func(index int, e *goquery.Selection) {

		item := ParseListItems(
			index,
			e,
			&selections,
			&itemParams.Query,
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

	h.logger.Info(ctx, fmt.Sprintf("Найдено %d вакансий", len(items)))

	return items, nil
}

func (h *Parser) UpdateVacancies(ctx context.Context, items []*models.ItemsModel) error {
	for _, item := range items {
		vacancy := models.Vacancy{
			Url:         item.Url,
			Title:       item.Title,
			Company:     item.Company,
			Salary:      item.Salary,
			Location:    item.Location,
			Experiences: item.Experience,
		}

		_, err := h.repo.AddVacancies(ctx, vacancy)
		if err != nil {
			h.logger.Error(ctx, fmt.Errorf("Ошибка при добавлении вакансии: %v", err))
			err = fmt.Errorf("ошибка добавления вакансии: %w", err)
			return err
		}
	}

	h.logger.Info(ctx, fmt.Sprintf("Добавлено %d вакансий в БД", len(items)))

	return nil
}
