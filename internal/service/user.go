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

//check is exxist by email - else craete new user, by email pwd, new uuid; open  new account
func (s *UserService) IsExistUser(id int) (string, error) {
	return s.userRepo.IsExistUser(id)
}

func (a *UserService) CreateUser(user *model.User) (int64, error) {

	// id, err := uuid.NewV4()
	// if err != nil {
	// 	return 0, err
	// }
	// log.Print(id.String(), "new uuid")

	// user.UUID = id.String()

	return a.userRepo.CreateUser(user)

}
