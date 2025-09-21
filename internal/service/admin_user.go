package service

import (
	"context"
	"eliborate/internal/repository"
	"eliborate/internal/service/validation"

	"golang.org/x/crypto/bcrypt"
)

type adminUserService struct {
	repo repository.AdminUserRepo
}

func InitAdminUserService(repo repository.AdminUserRepo) AdminUserService {
	return adminUserService{
		repo: repo,
	}
}

func (u adminUserService) UpdatePassword(ctx context.Context, id int, password string) error {
	if err := validation.ValidateID(id); err != nil {
		return err
	}
	if err := validation.ValidatePassword(password); err != nil {
		return err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := u.repo.UpdatePassword(ctx, id, string(hashedPass)); err != nil {
		return err
	}

	return nil
}
