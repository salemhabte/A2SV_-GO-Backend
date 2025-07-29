package usecases

import (
	"errors"
	"fmt"
	"go_clean/Domain"
	"go_clean/Infrastructure"
)

type UserUseCase struct{
	userRepo UserRepositoryInterface

}

func NewUserusecase( usrepo UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{
		userRepo: usrepo,
	}
}

func (uc *UserUseCase) Register(user *Domain.User) error {
	hash, err := Infrastructure.Hashpassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return uc.userRepo.CreateUser(user)
}

func (uc *UserUseCase) Login(username, password string) (string, error) {
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}
	if Infrastructure.ComparePassword(user.Password, password) != nil {
		return "", errors.New("invalid credentials")
	}
	fmt.Print(user.Role)
	return Infrastructure.GenerateToken(user.ID, user.Role)
}
