package usecases

import (
	"context"
	"errors"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/infrastructure"
)

type UserUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *UserUsecase {
	return &UserUsecase{
		repository: userRepository,
	}
}

func (userUsecase *UserUsecase) Register(user domain.User, cxt context.Context) (interface{}, error) {
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil{
		return nil, err
	}
	user.Password = hashedPassword

	count, err := userUsecase.repository.CountUserByEmail(user.Email, context.TODO())
	if err != nil{
		return nil, err
	}

	if count > 0{
		return nil, errors.New("Email already exists")
	}

	
	count, err = userUsecase.repository.CountUserByUsername(user.UserName, context.TODO())
	if err != nil{
		return nil, err
	}

	if count > 0{
		return nil, errors.New("Username already exists")
	}


	return userUsecase.repository.Register(user, context.TODO())
}

func (userUsecase *UserUsecase) Login(username, password string, cxt context.Context) (string, error) {
	user, err := userUsecase.repository.GetUserByUsername(username, context.TODO())
	if err != nil{
		return "", err
	}

	if !infrastructure.ComparePassword(password, user.Password){
		return "", errors.New("Incorrect username or password")
	}

	return infrastructure.GenerateToken(user)
}

func (userUsecase *UserUsecase) PromoteUser(username string, cxt context.Context) error {
	user, err := userUsecase.repository.GetUserByUsername(username, context.TODO())
	if err != nil{
		return err
	}

	if user.IsAdmin{
		return errors.New("the user with the given username is already an admin")
	}

	
	return userUsecase.repository.PromoteUser(username, context.TODO())
}
