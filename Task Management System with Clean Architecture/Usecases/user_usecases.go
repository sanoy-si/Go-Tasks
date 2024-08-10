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

func (uu *UserUsecase) Register(user domain.User) (interface{}, error) {
	count, err := uu.repository.CountUserByEmail(user.Email, context.TODO())
	if err != nil{
		return nil, err
	}
	
	if count > 0{
		return nil, errors.New("email already exists")
	}
	
	
	count, err = uu.repository.CountUserByUsername(user.UserName, context.TODO())
	if err != nil{
		return nil, err
	}
	
	if count > 0{
		return nil, errors.New("username already exists")
	}


	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil{
		return nil, err
	}
	user.Password = hashedPassword

	user.ID = infrastructure.GenerateID()
	
	count, err = uu.repository.CountAllUsers(context.TODO())
	if err != nil{
		return nil, err
	}
	
	if count == 0{
		user.IsAdmin = true
	}


	return uu.repository.Register(user, context.TODO())
}

func (uu *UserUsecase) Login(userCredentials domain.UserCredentials) (string, error) {
	if err := infrastructure.ValidateUserCredentials(&userCredentials); err != nil{
		return "", err
	}

	user, err := uu.repository.GetUserByUsername(userCredentials.UserName, context.TODO())
	if err != nil{
		return "", err
	}

	if !infrastructure.ComparePassword(userCredentials.Password, user.Password){
		return "", errors.New("incorrect username or password")
	}

	return infrastructure.GenerateToken(user)
}

func (uu *UserUsecase) PromoteUser(username string) error {
	user, err := uu.repository.GetUserByUsername(username, context.TODO())
	if err != nil{
		return err
	}

	if user.IsAdmin{
		return errors.New("the user with the given username is already an admin")
	}

	
	return uu.repository.PromoteUser(username, context.TODO())
}
