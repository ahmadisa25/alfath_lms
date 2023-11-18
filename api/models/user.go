package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `validate:"required" bson:"name,omitempty"`
	Password    string             `validate:"required" bson:"password,omitempty"`
	Email       string             `validate:"required,email" bson:"email,omitempty"`
	MobilePhone string             `validate:"required,numeric" bson:"mobilephone,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `bson:"updatedAt,omitempty"`
	Role        Role               `validate:"omitempty" bson:"role,omitempty"`
}
