package service

import (
	"context"
	"eliborate/internal/repository"
)

type categoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) CategoryService {
	return categoryService{
		repo: repo,
	}
}

func (c categoryService) Create(ctx context.Context, categoryName string) error {
	return c.repo.Create(ctx, categoryName)
}

func (c categoryService) GetAll(ctx context.Context) ([]string, error) {
	return c.repo.GetAll(ctx)
}

func (c categoryService) Update(ctx context.Context, id int, newName string) error {
	return c.repo.Update(ctx, id, newName)
}

func (c categoryService) Delete(ctx context.Context, name string) error {
	return c.repo.Delete(ctx, name)
}
