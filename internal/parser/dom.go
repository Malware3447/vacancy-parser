package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
	"vacancy-parser/internal/models"
)

func ParseListItems(index int, e *goquery.Selection, selection *models.ItemsList, query *models.ListParamsQuery, errorMap map[int]error) *models.ItemsModel {
	item := new(models.ItemsModel)

	selection.Title = e.Find(query.Title)
	if selection.Title == nil {
		err := fmt.Errorf("ошибка при получении заголовка элемента")
		errorMap[index] = err
		return nil
	} else {
		title := strings.TrimSpace(selection.Title.Text())
		item.Title = title
	}

	selection.Salary = e.Find(query.Salary)
	if selection.Salary == nil {
		item.Salary = -1
	} else {
		salaryText := strings.TrimSpace(selection.Salary.Text())
		salary, _ := strconv.Atoi(salaryText)
		item.Salary = int32(salary)
	}

	selection.Company = e.Find(query.Company)
	if selection.Company == nil {
		err := fmt.Errorf("ошибка при получении названия компании")
		errorMap[index] = err
		return nil
	} else {
		company := strings.TrimSpace(selection.Company.Text())
		item.Company = company
	}

	selection.Location = e.Find(query.Location)
	if selection.Location == nil {
		err := fmt.Errorf("ошибка при получении местоположения")
		errorMap[index] = err
		return nil
	} else {
		location := strings.TrimSpace(selection.Location.Text())
		item.Location = location
	}

	selection.Experience = e.Find(query.Experience)
	if selection.Experience == nil {
		err := fmt.Errorf("ошибка при получении опыта работы")
		errorMap[index] = err
		return nil
	} else {
		experience := strings.TrimSpace(selection.Experience.Text())
		item.Experience = experience
	}

	item.Url = query.Link

	return item
}
