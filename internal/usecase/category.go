package usecase

import (
	"context"

	"github.com/Akmyrat03/shop/internal/domain"
	"github.com/Akmyrat03/shop/internal/repository/postgres"
)

var _ CategoryUC = (*categoryUC)(nil)

type CategoryUC interface {
	Create(ctx context.Context, category domain.Category) (int, error)
}

type categoryUC struct {
	repo postgres.CategoryRepository
}

func NewCategoryUC(repo postgres.CategoryRepository) *categoryUC {
	return &categoryUC{repo: repo}
}

func (c *categoryUC) Create(ctx context.Context, category domain.Category) (int, error) {
	id, err := c.repo.Create(ctx, category)
	if err != nil {
		return 0, err
	}

	return id, nil
}
