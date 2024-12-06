package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/dto"
	"eliborate/internal/repository"
	"eliborate/internal/validators"
	"fmt"

	"eliborate/pkg/logging"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo   repository.UserRepo
	logger *logging.Log
}

func InitUserService(repo repository.UserRepo, logger *logging.Log) UserService {
	return userService{
		repo:   repo,
		logger: logger,
	}
}

func (u userService) Create(ctx context.Context, user dto.UserCreate) (int, error) {
	userConv := convertors.ToDomainUserCreate(user)

	if validErr := validators.IsPasswordValid(user.Password); validErr != nil {
		return 0, validErr
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userConv.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userConv.Password = string(hashedPass)

	id, err := u.repo.Create(ctx, userConv)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user create error '%s'", err.Error()))
		return 0, err
	}

	return id, nil
}

func (u userService) UpdatePassword(ctx context.Context, id int, password string) error {
	if validErr := validators.IsPasswordValid(password); validErr != nil {
		return validErr
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := u.repo.UpdatePassword(ctx, id, string(hashedPass)); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user update error '%s'", err.Error()))
		return err
	}

	return nil
}

func (u userService) Delete(ctx context.Context, id int) error {
	if err := u.repo.Delete(ctx, id); err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("user delete error '%s'", err.Error()))
		return err
	}

	return nil
}
