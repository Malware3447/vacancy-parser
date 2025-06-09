package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"vacancy-parser/internal/models"
)

func splitAndTakeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return s[:len(s)/2]
}

func ParseListItems(index int, e *goquery.Selection, selection *models.ItemsList, query *models.ListParamsQuery, errorMap map[int]error) *models.ItemsModel {
	item := new(models.ItemsModel)

	selection.Link = e.Find(query.Link)
	if selection.Link.Length() == 0 {
		err := fmt.Errorf("ошибка при получении ссылки на вакансию")
		errorMap[index] = err
		return nil
	} else {
		href, exists := selection.Link.Attr("href")
		if !exists {
			err := fmt.Errorf("атрибут href не найден")
			errorMap[index] = err
			return nil
		}
		item.Url = href
	}

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
	if selection.Salary == nil || selection.Salary.Length() == 0 {
		item.Salary = "не указано"
	} else {
		salary := strings.TrimSpace(selection.Salary.Text())
		salary = splitAndTakeFirst(salary)
		item.Salary = salary
	}

	selection.Company = e.Find(query.Company)
	if selection.Company == nil {
		err := fmt.Errorf("ошибка при получении названия компании")
		errorMap[index] = err
		return nil
	} else {
		company := strings.TrimSpace(selection.Company.Text())
		company = splitAndTakeFirst(company)
		item.Company = company
	}

	selection.Location = e.Find(query.Location)
	if selection.Location == nil {
		err := fmt.Errorf("ошибка при получении местоположения")
		errorMap[index] = err
		return nil
	} else {
		location := strings.TrimSpace(selection.Location.Text())
		location = splitAndTakeFirst(location)
		location = splitAndTakeFirst(location)
		item.Location = location
	}

	selection.Experience = e.Find(query.Experience)
	if selection.Experience == nil {
		err := fmt.Errorf("ошибка при получении опыта работы")
		errorMap[index] = err
		return nil
	} else {
		experience := strings.TrimSpace(selection.Experience.Text())
		experience = splitAndTakeFirst(experience)
		item.Experience = experience
	}

	return item
}
