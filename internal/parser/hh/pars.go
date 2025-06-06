package hh

import (
	"context"
	"github.com/x3a-tech/logit-go"
)

type HhParser struct {
	Logger *logit.Logger
}

type HhParams struct {
	Logger *logit.Logger
}

func NewHhParser(ctx context.Context, params HhParams) *HhParser {
	return &HhParser{
		Logger: params.Logger,
	}
}
