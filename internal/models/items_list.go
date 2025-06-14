package models

import "github.com/PuerkitoBio/goquery"

type ItemParams struct {
	BaseUrl string
	Query   ListParamsQuery
}

type ListParamsQuery struct {
	Items      string
	Link       string
	Title      string
	Salary     string
	Company    string
	Location   string
	Experience string
}

type ItemsList struct {
	Items      *goquery.Selection
	Link       *goquery.Selection
	Title      *goquery.Selection
	Salary     *goquery.Selection
	Company    *goquery.Selection
	Location   *goquery.Selection
	Experience *goquery.Selection
}
