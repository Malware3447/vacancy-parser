package pg

import (
	"context"
	"vacancy-parser/internal/models"
)

type Repository interface {
	AddVacancies(ctx context.Context, params models.Vacancy) (id int32, err error)
}
