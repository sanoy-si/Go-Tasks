package data

import (
	"Task_Management_System_with_JWT/models"
	"context"
	"errors"

	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)


type UserService interface{
	Register(user models.User) (interface{}, error)
	Login(username, password string) (string, error)
	PromoteUser(username string) error
}

type MongoDBUserService struct{
	client *mongo.Client
}


func NewMongoDBUserService() (*MongoDBUserService, error){
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		return &MongoDBUserService{}, err
	}

	return &MongoDBUserService{
		client: client,
	}, nil
}

func (service *MongoDBUserService) Register(user models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return nil, err
	}
	user.Password = string(hashedPassword)


	collection := service.client.Database("task-manager").Collection("users")
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


func (service *MongoDBUserService) Login(username, password string) (string, error) {
	var user models.User
	collection := service.client.Database("task-manager").Collection("users")
	filter := bson.M{"username":username}

	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil{
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":user.UserName,
		"is_admin":user.IsAdmin,
		"expires_at":time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil{
		return "", err
	}

	return tokenString, nil

}


func (service *MongoDBUserService) PromoteUser(username string) error{
	var user models.User
	collection := service.client.Database("task-manager").Collection("users")
	filter := bson.M{"username": username}

	if err := collection.FindOne(context.TODO(), filter).Decode(&user); err != nil{
		if err.Error() == "mongo: no documents in result"{
			return errors.New("user not foud")
		}
		return err
	}

	if user.IsAdmin{
		return errors.New("user is already an admin")
	}

	update := bson.M{
		"$set":bson.M{"is_admin": true},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil{
		return err
	}

	return nil

}

