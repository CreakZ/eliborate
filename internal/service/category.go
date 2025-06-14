package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
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

func (c categoryService) GetAll(ctx context.Context) ([]domain.Category, error) {
	categories, err := c.repo.GetAll(ctx)
	if err != nil {
		return []domain.Category{}, err
	}
	return convertors.EntityCategoriesToDomain(categories), nil
}

func (c categoryService) Update(ctx context.Context, id int, newName string) error {
	return c.repo.Update(ctx, id, newName)
}

func (c categoryService) Delete(ctx context.Context, id int) error {
	return c.repo.Delete(ctx, id)
}
