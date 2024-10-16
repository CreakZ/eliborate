package service

import (
	"context"
	"yurii-lib/internal/repository"
	"yurii-lib/pkg/lgr"
)

type publicService struct {
	repo   repository.PublicRepo
	logger *lgr.Log
}

func InitPublicService(repo repository.PublicRepo, logger *lgr.Log) PublicService {
	return publicService{
		repo:   repo,
		logger: logger,
	}
}

func (p publicService) GetByLogin(ctx context.Context, userType, login string) (int, string, error) {
	id, password, err := p.repo.GetByLogin(ctx, userType, login)
	if err != nil {
		return 0, "", err
	}

	return id, password, nil
}
