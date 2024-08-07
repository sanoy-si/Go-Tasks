package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `json:"first_name" validate:"required,min=2,max=100"`
	LastName  string             `json:"last_name" validate:"required,min=2,max=100"`
	Email     string             `json:"email" validate:"email,required"`
	UserName  string			 `json:"username" validate:"required,min=5"`
	Password  string             `json:"password" validate:"required,min=6"`
	IsAdmin   bool               `json:"is_admin"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
