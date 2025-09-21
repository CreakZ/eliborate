package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
	"eliborate/internal/repository"
	"eliborate/internal/service/validation"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo repository.UserRepo
}

func InitUserService(repo repository.UserRepo) UserService {
	return userService{
		repo: repo,
	}
}

func (u userService) Create(ctx context.Context, user domain.UserCreate) (int, error) {
	if err := validation.ValidateUserCreate(user); err != nil {
		return 0, err
	}

	userEntity := convertors.DomainUserCreateToEntity(user)

	if err := validation.ValidatePassword(user.Password); err != nil {
		return 0, err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userEntity.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	userEntity.Password = string(hashedPass)

	id, err := u.repo.Create(ctx, userEntity)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u userService) UpdatePassword(ctx context.Context, id int, password string) error {
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

func (u userService) Delete(ctx context.Context, id int) error {
	if err := validation.ValidateID(id); err != nil {
		return err
	}

	if err := u.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
