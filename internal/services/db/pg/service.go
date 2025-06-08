package pg

import (
	"context"
	"vacancy-parser/internal/db/repository/pg"
	"vacancy-parser/internal/models"
)

type Service struct {
	repo pg.Repository
}

func NewService(repo pg.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddVacancies(ctx context.Context, params models.Vacancy) (id int32, err error) {
	return s.repo.AddVacancies(ctx, params)
}
