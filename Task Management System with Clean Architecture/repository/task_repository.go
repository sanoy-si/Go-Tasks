package repositroy

import (
	"context"
	"errors"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type TaskRepository struct{
	database *mongo.Database
	collection string
}

func NewTaskRepository(db *mongo.Database, collection string) *TaskRepository{
	return &TaskRepository{
		database: db,
		collection: collection,
	}
}

func (tr *TaskRepository) GetTasks(ctx context.Context) ([]domain.Task, error) {
	collection := tr.database.Collection(tr.collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []domain.Task{}, err
	}

	allTasks := []domain.Task{}

	for cursor.Next(ctx) {
		var task domain.Task

		err := cursor.Decode(&task)
		if err != nil {
			return []domain.Task{}, err
		}

		allTasks = append(allTasks, task)

	}

	return allTasks, nil
}


func (tr *TaskRepository) GetTaskByID(id string, ctx context.Context)(domain.Task, error){
	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"_id":id}
	
	var task domain.Task
	err := collection.FindOne(ctx, filter).Decode(&task)
	if err != nil{
		return domain.Task{}, err
	}

	return task, nil
}

func (tr *TaskRepository) CreateTask(newTask domain.Task, ctx context.Context) (domain.Task, error){
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(ctx, newTask)

	if err != nil{
		return domain.Task{}, err
	}
	
	return newTask, nil
}

func (tr *TaskRepository) UpdateTask(id string, updatedTask domain.Task, ctx context.Context) (domain.Task, error){
	collection := tr.database.Collection(tr.collection)

	_, err := tr.GetTaskByID(id, ctx)
	if err != nil{
		return domain.Task{}, err
		} 
		
	filter := bson.M{"__id": id}
	update := bson.M{"$set": bson.M{
		"title": updatedTask.Title,
		"description": updatedTask.Description,
		"due_date": updatedTask.DueDate,
		"completed": updatedTask.Completed,
	}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil{
		return domain.Task{}, err
	}

	return tr.GetTaskByID(id, ctx)
}


func (tr *TaskRepository) DeleteTask(id string, ctx context.Context) error{
	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)

	if err != nil{
		return err
	}

	if deleteResult.DeletedCount == 0{
		return errors.New("task not found")
	}

	return nil
}
