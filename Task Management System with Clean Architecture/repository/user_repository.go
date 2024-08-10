package repositroy

import (
	"context"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct{
	database mongo.Database
	collection string
} 

func NewUserRepository(db mongo.Database, collection string) *userRepository{
	return &userRepository{
		database: db,
		collection: collection,
	}
}

func (ur *userRepository) Register(user domain.User, ctx context.Context) (interface{}, error) {

	collection := ur.database.Collection(ur.collection)
	insertedId, err := collection.InsertOne(context.TODO(), user)
	if err != nil{
		return nil, err
	}
	
	return insertedId, nil
}


func (ur *userRepository) PromoteUser(username string) error{
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username": username}
	update := bson.M{
		"$set":bson.M{"is_admin": true},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil{
		return err
	}

	return nil
}

