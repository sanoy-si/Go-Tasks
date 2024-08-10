package usecases

import (
	"context"
	"errors"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/infrastructure"
)

type userUsecase struct {
	repository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{
		repository: userRepository,
	}
}

func (uu *userUsecase) Register(user domain.User, cxt context.Context) (interface{}, error) {
	if err := infrastructure.ValidateUser(&user); err != nil{
		return nil, err
	}

	
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
	
	count, err = uu.repository.CountAllUsers()
	if err != nil{
		return nil, err
	}
	
	if count == 0{
		user.IsAdmin = true
	}


	return uu.repository.Register(user, context.TODO())
}

func (uu *userUsecase) Login(userCredentials domain.UserCredentials, cxt context.Context) (string, error) {
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

func (uu *userUsecase) PromoteUser(username string, cxt context.Context) error {
	user, err := uu.repository.GetUserByUsername(username, context.TODO())
	if err != nil{
		return err
	}

	if user.IsAdmin{
		return errors.New("the user with the given username is already an admin")
	}

	
	return uu.repository.PromoteUser(username, context.TODO())
}
