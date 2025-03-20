package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
	"eliborate/internal/repository"
	"eliborate/pkg/logging"
)

type publicService struct {
	repo   repository.PublicRepo
	logger *logging.Log
}

func InitPublicService(repo repository.PublicRepo, logger *logging.Log) PublicService {
	return publicService{
		repo:   repo,
		logger: logger,
	}
}

func (p publicService) GetUserByLogin(ctx context.Context, login string) (domain.User, error) {
	user, err := p.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return domain.User{}, err
	}
	return convertors.EntityUserToDomain(user), nil
}

func (p publicService) GetAdminUserByLogin(ctx context.Context, login string) (domain.AdminUser, error) {
	user, err := p.repo.GetAdminUserByLogin(ctx, login)
	if err != nil {
		return domain.AdminUser{}, err
	}
	return convertors.EntityAdminUserToDomain(user), nil
}
