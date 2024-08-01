package service

import (
	"context"
	"fmt"
	"yurii-lib/internal/convertors"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/repository"

	"yurii-lib/pkg/log"
)

type userService struct {
	repo   repository.UserRepo
	logger *log.Log
}

func InitUserService(repo repository.UserRepo, logger *log.Log) UserService {
	return userService{
		repo:   repo,
		logger: logger,
	}
}

func (u userService) CreateAdminUser(ctx context.Context, user dto.AdminUserCreate) (int, error) {
	userConv := convertors.ToDomainAdminUserCreate(user)

	id, err := u.repo.CreateAdminUser(ctx, userConv)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user create error '%s'", err.Error()))
		return 0, err
	}

	return id, nil
}

func (u userService) CreateUser(ctx context.Context, user dto.UserCreate) (int, error) {
	userConv := convertors.ToDomainUserCreate(user)

	id, err := u.repo.CreateUser(ctx, userConv)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user create error '%s'", err.Error()))
		return 0, err
	}

	return id, nil
}

func (u userService) GetAdminUserPassword(ctx context.Context, id int) (string, error) {
	password, err := u.repo.GetAdminUserPassword(ctx, id)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get admin user password error '%s'", err.Error()))
		return "", err
	}

	return password, nil
}

func (u userService) GetUserPassword(ctx context.Context, id int) (string, error) {
	password, err := u.repo.GetUserPassword(ctx, id)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get user password error '%s'", err.Error()))
		return "", err
	}

	return password, nil
}

func (u userService) UpdateAdminUserPassword(ctx context.Context, id int, password string) error {
	if err := u.repo.UpdateAdminUserPassword(ctx, id, password); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user update error '%s'", err.Error()))
		return err
	}

	return nil
}

func (u userService) UpdateUserPassword(ctx context.Context, id int, password string) error {
	if err := u.repo.UpdateUserPassword(ctx, id, password); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user update error '%s'", err.Error()))
		return err
	}

	return nil
}

func (u userService) DeleteAdminUser(ctx context.Context, id int) error {
	if err := u.repo.DeleteAdminUser(ctx, id); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user delete error '%s'", err.Error()))
		return err
	}

	return nil
}

func (u userService) DeleteUser(ctx context.Context, id int) error {
	if err := u.repo.DeleteUser(ctx, id); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user delete error '%s'", err.Error()))
		return err
	}

	return nil
}
