package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
)

var validate = validator.New()

func ValidateUser(user *domain.User) error{
	if err := validate.Struct(user); err != nil{
		return err
	}

	return nil
}

func ValidateTask(task *domain.Task) error{
	if err := validate.Struct(task); err != nil{
		return err
	}

	return nil
}

func ValidateUserCredentials(userCredentials *domain.UserCredentials) error{
	if err := validate.Struct(userCredentials); err != nil{
		return err
	}

	return nil
}