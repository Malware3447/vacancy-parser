package pg

import (
	"context"
	"vacancy-parser/internal/models"
)

type Repository interface {
	AddVacancy(ctx context.Context, params models.Vacancy) (id int32, err error)
}
