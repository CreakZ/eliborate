package service

import (
	"context"
	"fmt"
	"yurii-lib/internal/repository"
	"yurii-lib/pkg/lgr"

	"golang.org/x/crypto/bcrypt"
)

type adminUserService struct {
	repo   repository.AdminUserRepo
	logger *lgr.Log
}

func InitAdminUserService(repo repository.AdminUserRepo, logger *lgr.Log) AdminUserService {
	return adminUserService{
		repo:   repo,
		logger: logger,
	}
}

/*
func (u adminUserService) Create(ctx context.Context, user dto.AdminUserCreate) (int, error) {
	userConv := convertors.ToDomainAdminUserCreate(user)

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userConv.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userConv.Password = string(hashedPass)

	id, err := u.repo.Create(ctx, userConv)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user create error '%s'", err.Error()))
		return 0, err
	}

	return id, nil
}
*/

func (u adminUserService) GetPassword(ctx context.Context, id int) (string, error) {
	password, err := u.repo.GetPassword(ctx, id)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get admin user password error '%s'", err.Error()))
		return "", err
	}

	return password, nil
}

func (u adminUserService) UpdatePassword(ctx context.Context, id int, password string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := u.repo.UpdatePassword(ctx, id, string(hashedPass)); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user update error '%s'", err.Error()))
		return err
	}

	return nil
}

/*
func (u adminUserService) Delete(ctx context.Context, id int) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("admin user delete error '%s'", err.Error()))
		return err
	}

	return nil
}
*/
