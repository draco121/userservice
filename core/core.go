package core

import (
	"context"
	"userservice/repository"

	"github.com/draco121/common/models"

	"github.com/draco121/common/utils"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
}

type userService struct {
	IUserService
	repo repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	us := userService{
		repo: repository,
	}
	return us
}

func (s userService) CreateUser(ctx context.Context, user *models.User) (string, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	} else {
		user.Password = hashedPassword
		id, err := s.repo.InsertOne(ctx, user)
		if err != nil {
			return "", err
		} else {
			return id, nil
		}
	}
}

func (s userService) GetUserById(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.FindOneById(ctx, id)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (s userService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repo.FindOneByEmail(ctx, email)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (s userService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	newPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	} else {
		user.Password = newPassword
		user, err := s.repo.UpdateOne(ctx, user)
		if err != nil {
			return nil, err
		} else {
			return user, nil
		}
	}
}

func (s userService) DeleteUser(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.DeleteOneById(ctx, id)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
