package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/dto"
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

func (p publicService) GetUserByLogin(ctx context.Context, login string) (dto.User, error) {
	user, err := p.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return dto.User{}, err
	}
	return convertors.ToDtoUser(user), nil
}

func (p publicService) GetAdminUserByLogin(ctx context.Context, login string) (dto.AdminUser, error) {
	user, err := p.repo.GetAdminUserByLogin(ctx, login)
	if err != nil {
		return dto.AdminUser{}, err
	}
	return convertors.ToDtoAdminUser(user), nil
}
