package pg

import (
	"context"
	"vacancy-parser/internal/db/repository/pg"
	"vacancy-parser/internal/structs/vacancy"
)

type Service struct {
	repo pg.Repository
}

func NewService(repo pg.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddVacancy(ctx context.Context, params vacancy.Vacancy) (id int32, err error) {
	return s.repo.AddVacancy(ctx, params)
}
