package domain

import (
	"time"
	"context"
)

type Task struct {
	ID          string    `bson:"_id"`
	Title       string    `json:"title" validate:"required,min=2,max=30"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
}

type User struct {
	ID        string			 `bson:"_id"`
	FirstName string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=100"`
	LastName  string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=100"`
	Email     string             `bson:"email" json:"email" validate:"email,required"`
	UserName  string			 `bson:"username" json:"username" validate:"required,min=5"`
	Password  string             `bson:"password" json:"password" validate:"required,min=6"`
	IsAdmin   bool               `bson:"is_admin"`
}

type UserCredentials struct{
	UserName  string			 `json:"username" validate:"required"`
	Password  string             `json:"password" validate:"required"`
}

type TaskUsecase interface{
	GetTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	CreateTask(newTask Task) (Task, error)
	UpdateTask(id string, updatedTask Task) (Task, error)
	DeleteTask(id string) error
}


type UserUsecase interface{
	Register(user User) (interface{}, error)
	Login(userCredentials UserCredentials) (string, error)
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
	PromoteUser(username string, ctx context.Context) error
	CountUserByUsername(username string, ctx context.Context) (int, error)
	CountUserByEmail(email string, ctx context.Context) (int, error)
	CountAllUsers()(int, error)
	GetUserByUsername(usename string, ctx context.Context) (User, error)
	
}