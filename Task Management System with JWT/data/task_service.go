package data

import (
	"Task_Management_System_with_JWT/models"
	"errors"
	"fmt"
	"strconv"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PersistentTaskManagementService struct{
	client *mongo.Client

}

func NewPersistentTaskManagementService() *PersistentTaskManagementService{
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		log.Fatal("Couldn't Connect to the databse.", err)
	}

	return &PersistentTaskManagementService{
		client: client,
	}
}


func (service *PersistentTaskManagementService) GetTasks() []models.Task{
	collection := service.client.Database("task_manager").Collection("tasks")
	
	cursor, err := collection.Find(context.TODO(), bson.M{})
	
	if err != nil{
		defer log.Fatal(err)
		return []models.Task{}
	}
	
	allTasks := []models.Task{}

	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task)

		if err != nil{
			defer log.Fatal(err)
			return []models.Task{}
		}
		
		allTasks = append(allTasks, task)

	}

	return allTasks
}

func (service *PersistentTaskManagementService) GetTask(id string) (models.Task, error){
	collection := service.client.Database("task_manager").Collection("tasks")
	filter := bson.M{"id":id}
	
	var task models.Task
	err := collection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil{
		return models.Task{}, err
	}

	return task, nil
}

func (service *PersistentTaskManagementService) CreateTask(newTask models.Task) models.Task{
	collection := service.client.Database("task_manager").Collection("tasks")
	insertResult, err := collection.InsertOne(context.TODO(), newTask)

	if err != nil{
		log.Fatal(err)
	}
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil{
		log.Fatal(err)
	}
	filter := bson.M{"_id": insertResult.InsertedID}
	update := bson.M{"$set": bson.M{"id": strconv.Itoa(int(count))}}
	
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		log.Fatal(err)
	}
	
	newTask.ID = strconv.Itoa(int(count))
	return newTask
}

func (service *PersistentTaskManagementService) UpdateTask(id string, updatedTask models.Task) (models.Task, error){
	collection := service.client.Database("task_manager").Collection("tasks")
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{
		"id": id,
		"title": updatedTask.Title,
		"description": updatedTask.Description,
		"due_date": updatedTask.DueDate,
		"status": updatedTask.Status,
	}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil{
		return models.Task{}, err
	}

	updatedTask.ID = id
	return updatedTask, nil
}

func (service *PersistentTaskManagementService) DeleteTask(id string) error{
	collection := service.client.Database("task_manager").Collection("tasks")
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

