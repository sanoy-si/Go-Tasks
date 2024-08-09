package domain

import (
	"time"
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
