package service

import (
	"context"
	"fmt"
	"yurii-lib/internal/convertors"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/repository"
	"yurii-lib/pkg/log"
)

type adminUserService struct {
	repo   repository.AdminUserRepo
	logger *log.Log
}

func InitAdminUserService(repo repository.AdminUserRepo, logger *log.Log) AdminUserService {
	return adminUserService{
		repo:   repo,
		logger: logger,
	}
}

func (u adminUserService) CreateAdminUser(ctx context.Context, user dto.AdminUserCreate) (int, error) {
	userConv := convertors.ToDomainAdminUserCreate(user)

	id, err := u.repo.CreateAdminUser(ctx, userConv)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user create error '%s'", err.Error()))
		return 0, err
	}

	return id, nil
}

func (u adminUserService) GetAdminUserPassword(ctx context.Context, id int) (string, error) {
	password, err := u.repo.GetAdminUserPassword(ctx, id)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get admin user password error '%s'", err.Error()))
		return "", err
	}

	return password, nil
}

func (u adminUserService) UpdateAdminUserPassword(ctx context.Context, id int, password string) error {
	if err := u.repo.UpdateAdminUserPassword(ctx, id, password); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user update error '%s'", err.Error()))
		return err
	}

	return nil
}

func (u adminUserService) DeleteAdminUser(ctx context.Context, id int) error {
	if err := u.repo.DeleteAdminUser(ctx, id); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user delete error '%s'", err.Error()))
		return err
	}

	return nil
}
