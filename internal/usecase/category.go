package usecase

import (
	"context"

	"github.com/Akmyrat03/shop/internal/database"
	"github.com/Akmyrat03/shop/internal/domain"
)

var _ CategoryUC = (*categoryUC)(nil)

type CategoryUC interface {
	Create(ctx context.Context, category domain.Category) (int, error)
}

type categoryUC struct {
	repo database.PSQLDBStore
}

func NewCategoryUC(repo database.PSQLDBStore) *categoryUC {
	return &categoryUC{repo: repo}
}

func (c *categoryUC) Create(ctx context.Context, category domain.Category) (int, error) {
	var err error
	var categoryID int

	err = c.repo.WithTransaction(ctx, func(dataStore database.PSQLDBStore) error {
		id, err := dataStore.CategoryRepo().Create(ctx, category)
		if err != nil {
			return 0, err
		}

		err = dataStore.TagRepo().Create(ctx)
		if err != nil {
			return err
		}

		categoryID = id

		return nil
	})
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}
