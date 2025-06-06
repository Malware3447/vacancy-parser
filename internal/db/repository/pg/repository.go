package pg

import (
	"context"
	"vacancy-parser/internal/structs/vacancy"
)

type Repository interface {
	AddVacancy(ctx context.Context, params vacancy.Vacancy) (id int32, err error)
}
