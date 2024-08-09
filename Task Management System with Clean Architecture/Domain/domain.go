package domain

import (
	"time"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email     string             `bson:"email" json:"email" validate:"email,required"`
	UserName  string			 `bson:"username" json:"username" validate:"required,min=5"`
	Password  string             `bson:"password" json:"password" validate:"required,min=6"`
	IsAdmin   bool               `bson:"is_admin" json:"is_admin"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

type TaskUsecase interface{
	GetTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	CreateTask(newTask Task) (Task, error)
	UpdateTask(id string, updatedTask Task) (Task, error)
	DeleteTask(id string) error
}


type UserUsecase interface{
	Register(user User) (interface{}, error)
	Login(username, password string) (string, error)
	PromoteUser(username string) error
}

type TaskRepository interface{
	GetTasks(ctx context.Context) ([]Task, error)
	GetTaskByID(id string, ctx context.Context) (Task, error)
	CreateTask(newTask Task, ctx context.Context) (Task, error)
	UpdateTask(id string, updatedTask Task, ctx context.Context) (Task, error)
	DeleteTask(id string, ctx context.Context) error
	
}

type UserRepository interface{
	Register(user User, ctx context.Context) (interface{}, error)
	Login(username, password string, ctx context.Context) (string, error)
	PromoteUser(username string, ctx context.Context) error
	CountUserByUsername(username string, ctx context.Context)
	CountUserByEmail(email string, ctx context.Context)
	GetUserByUsername(usename string, ctx context.Context)
	
}