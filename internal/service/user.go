package service

import (
	"context"
	"fmt"
	"yurii-lib/internal/convertors"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/repository"

	"yurii-lib/pkg/lgr"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo   repository.UserRepo
	logger *lgr.Log
}

func InitUserService(repo repository.UserRepo, logger *lgr.Log) UserService {
	return userService{
		repo:   repo,
		logger: logger,
	}
}

func (u userService) Create(ctx context.Context, user dto.UserCreate) (int, error) {
	userConv := convertors.ToDomainUserCreate(user)

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

func (u userService) CheckByLogin(ctx context.Context, login string) (bool, error) {
	exists, err := u.repo.CheckByLogin(ctx, login)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get user password error '%s'", err.Error()))
		return false, err
	}

	return exists, nil
}

func (u userService) GetPassword(ctx context.Context, id int) (string, error) {
	password, err := u.repo.GetPassword(ctx, id)
	if err != nil {
		u.logger.InfoLogger.Info().Msg(fmt.Sprintf("get user password error '%s'", err.Error()))
		return "", err
	}

	return password, nil
}

func (u userService) UpdatePassword(ctx context.Context, id int, password string) error {
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
