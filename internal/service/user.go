package service

import (
	"avito/internal"
	"avito/internal/model"
)

type UserService struct {
	userRepo internal.UserRepoInterface
}

func NewUserService(repo internal.UserRepoInterface) internal.UserServiceInterface {
	return &UserService{userRepo: repo}
}

func (s *UserService) IsExistUser(id int) (string, error) {
	return s.userRepo.IsExistUser(id)
}

func (a *UserService) CreateUser(user *model.User) (int64, error) {
	return a.userRepo.CreateUser(user)
}
