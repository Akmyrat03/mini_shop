package postgres

import (
	"context"

	"github.com/Akmyrat03/shop/internal/domain"
	"github.com/Akmyrat03/shop/pkg/connection"
)

var _ CategoryRepository = (*categoryRepository)(nil)

type CategoryRepository interface {
	Create(ctx context.Context, category domain.Category) (int, error)
}

type categoryRepository struct {
	psqlDB connection.DB
}

func NewCategoryRepository(psqlDB connection.DB) *categoryRepository {
	return &categoryRepository{psqlDB: psqlDB}
}

func (r *categoryRepository) Create(ctx context.Context, category domain.Category) (int, error) {
	var id int

	query := `INSERT INTO category (name_tk, name_en, name_ru) VALUES ($1, $2, $3) RETURNING id`

	err := r.psqlDB.QueryRow(ctx, query, category.NameTK, category.NameEN, category.NameRU).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
