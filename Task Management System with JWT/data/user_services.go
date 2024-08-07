package data

import (
	"Task_Management_System_with_JWT/models"
	"context"
	"errors"
	// "os"
	"time"

	"github.com/dgrijalva/jwt-go"
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


	filter = bson.M{"username":user.UserName}
	count, err = collection.CountDocuments(context.TODO(), filter)
	if err != nil{
		return nil, err
	}

	if count > 0{
		return nil, errors.New("username already exists")
	}


	if user.CreatedAt.IsZero(){
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = user.CreatedAt


	count, err = collection.CountDocuments(context.TODO(), bson.D{})
	if err != nil{
		return nil, err
	}
	
	if count == 0{
		user.IsAdmin = true
	}


	insertedId, err := collection.InsertOne(context.TODO(), user)
	if err != nil{
		return nil, err
	}
	

	return insertedId, nil

}


func (service *PersistentTaskManagementService) Login(userName, password string) (string, error) {
	var user models.User
	collection := service.client.Database("task_manager").Collection("users")
	filter := bson.M{"username":userName}

	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil{
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return "", err
	}

	secretKey := "my_secret_key"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":user.UserName,
		"is_admin":user.IsAdmin,
		"expires_at":time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil{
		return "", err
	}

	return tokenString, nil

}



