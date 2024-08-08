package data

import (
	"Task_Management_System_with_JWT/models"
	"context"
	"errors"

	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type PersistentTaskManagementService struct{
	client *mongo.Client

}



func NewPersistentTaskManagementService() (*PersistentTaskManagementService, error){
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		return &PersistentTaskManagementService{}, err
	}

	return &PersistentTaskManagementService{
		client: client,
	}, nil
}


func (service *PersistentTaskManagementService) GetTasks() ([]models.Task, error){
	collection := service.client.Database("task-manager").Collection("tasks")
	
	cursor, err := collection.Find(context.TODO(), bson.M{})
	
	if err != nil{
		return []models.Task{}, err
	}
	
	allTasks := []models.Task{}

	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task)

		if err != nil{
			return []models.Task{}, err
		}
		
		allTasks = append(allTasks, task)

	}

	return allTasks, nil
}

func (service *PersistentTaskManagementService) GetTask(id string) (models.Task, error){
	collection := service.client.Database("task-manager").Collection("tasks")
	filter := bson.M{"id":id}
	
	var task models.Task
	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil{
		return models.Task{}, err
	}

	return task, nil
}

func (service *PersistentTaskManagementService) CreateTask(newTask models.Task) (models.Task, error){
	collection := service.client.Database("task-manager").Collection("tasks")
	insertResult, err := collection.InsertOne(context.TODO(), newTask)

	if err != nil{
		return models.Task{}, err
	}
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil{
		return models.Task{}, err
	}
	filter := bson.M{"_id": insertResult.InsertedID}
	update := bson.M{"$set": bson.M{"id": strconv.Itoa(int(count))}}
	
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return models.Task{}, err
	}
	
	newTask.ID = strconv.Itoa(int(count))
	return newTask, nil
}

func (service *PersistentTaskManagementService) UpdateTask(id string, updatedTask models.Task) (models.Task, error){
	collection := service.client.Database("task-manager").Collection("tasks")
	filter := bson.M{"id": id}
	_, err := service.GetTask(id)
	if err != nil{
		return models.Task{}, err
	} 

	update := bson.M{"$set": bson.M{
		"title": updatedTask.Title,
		"description": updatedTask.Description,
		"due_date": updatedTask.DueDate,
		"status": updatedTask.Status,
	}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return models.Task{}, err
	}

	updatedTask.ID = id
	return updatedTask, nil
}

func (service *PersistentTaskManagementService) DeleteTask(id string) error{
	collection := service.client.Database("task-manager").Collection("tasks")
	filter := bson.M{"id": id}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)

	fmt.Println(deleteResult.DeletedCount)
	if err != nil{
		return err
	}

	if deleteResult.DeletedCount == 0{
		return errors.New("task not found")
	}

	return nil
}

