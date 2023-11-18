package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `validate:"required" bson:"name,omitempty"`
	CreatedAt time.Time          `bson:"createdat,omitempty"`
	CreatedBy string             `validate:"required" bson:"createdby,omitempty"`
	UpdatedAt time.Time          `bson:"updatedat,omitempty"`
	UpdatedBy string             `validate:"required" bson:"updatedby,omitempty"`
}
