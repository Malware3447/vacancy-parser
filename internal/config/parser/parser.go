package parser

import "vacancy-parser/internal/config/parser/website"

type Parser struct {
	Hh website.Website `yaml:"hh"`
}
