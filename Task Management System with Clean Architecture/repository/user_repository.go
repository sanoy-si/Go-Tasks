package repositroy

import (
	"context"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct{
	database *mongo.Database
	collection string
} 

func NewUserRepository(db *mongo.Database, collection string) *userRepository{
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


func (ur *userRepository) PromoteUser(username string, ctx context.Context) error{
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username": username}
	update := bson.M{
		"$set":bson.M{"is_admin": true},
	}
	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil{
		return err
	}

	return nil
}

func (ur *userRepository) CountAllUsers(ctx context.Context) (int, error){
	collection := ur.database.Collection(ur.collection)
	count, err := collection.CountDocuments(ctx, bson.M{})

	if err != nil{
		return 0, err
	}

	return int(count), nil
}


func (ur *userRepository) CountUserByEmail(email string, ctx context.Context) (int, error){
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"email":email}
	count, err := collection.CountDocuments(ctx, filter)

	if err != nil{
		return 0, err
	}

	return int(count), nil
}


func (ur *userRepository) CountUserByUsername(userName string, ctx context.Context) (int, error){
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username":userName}
	count, err := collection.CountDocuments(ctx, filter)

	if err != nil{
		return 0, err
	}

	return int(count), nil
}


func (ur *userRepository) GetUserByUsername(userName string, ctx context.Context) (domain.User, error){
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username":userName}
	var user domain.User
	
	if err := collection.FindOne(ctx, filter).Decode(&user); err != nil{
		return domain.User{}, err
	}

	return user, nil
}