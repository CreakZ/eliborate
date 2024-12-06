package service

import (
	"context"
	"eliborate/internal/repository"
	"eliborate/internal/validators"
	"eliborate/pkg/logging"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type adminUserService struct {
	repo   repository.AdminUserRepo
	logger *logging.Log
}

func InitAdminUserService(repo repository.AdminUserRepo, logger *logging.Log) AdminUserService {
	return adminUserService{
		repo:   repo,
		logger: logger,
	}
}

func (u adminUserService) UpdatePassword(ctx context.Context, id int, password string) error {
	if validErr := validators.IsPasswordValid(password); validErr != nil {
		return validErr
	}

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
