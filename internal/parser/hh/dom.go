package hh

import (
	"github.com/PuerkitoBio/goquery"
	"vacancy-parser/internal/models"
)

func ParseListItems(index int, selection *goquery.Selection, query *models.ListParamsQuery, errorMap map[int]error) *models.ItemsModel {
	item := new(models.ItemsModel)

	return item
}
