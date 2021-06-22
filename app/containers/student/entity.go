package student

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Student struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
	Email string `json:"email" bson:"email"`
	Gender string `json:"gender" bson:"gender"`
	IsActive bool `json:"is_active" bson:"is_active"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type FormatRequest struct {
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email string `json:"email"`
	Gender string `json:"gender"`
}

type FormatResponse struct {
	ID primitive.ObjectID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
