package data

import (
	"Task_Management_System_with_JWT/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)


func (service *PersistentTaskManagementService) Register(user models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	user.Password = string(hashedPassword)

	collection := service.client.Database("task_manager").Collection("users")
	filter := bson.M{"email":user.Email}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil{
		return nil, err
	}

	if count > 0{
		return nil, errors.New("email already exists")
	}

	if user.CreatedAt.IsZero(){
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = user.CreatedAt

	insertedId, err := collection.InsertOne(context.TODO(), user)
	if err != nil{
		return nil, err
	}

	return insertedId, nil

}

