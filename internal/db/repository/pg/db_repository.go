package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"vacancy-parser/internal/models"
)

type RepositoryPg struct {
	db *pgxpool.Pool
}

func NewRepositoryPg(db *pgxpool.Pool) Repository {
	return &RepositoryPg{db: db}
}

func (r *RepositoryPg) AddVacancy(ctx context.Context, params models.Vacancy) (id int32, err error) {
	const q = `
	INSERT INTO vacancy (link, title, salary, company, location, experience, source_id, created_at)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	created_at := time.Now()

	_, err = r.db.Exec(ctx, q, params.Url, params.Title, params.Salary, params.Company, params.Location, params.Experiences, params.Source_id, created_at)
	if err != nil {
		return 0, fmt.Errorf("failed to upsert vacancy: %v", err)
	}

	return 0, nil
}
