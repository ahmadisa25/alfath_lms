package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Announcement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `validate:"required" json:"title,omitempty"`
	Description string             `validate:"required" json:"description,omitempty"`
	FileUrl     string             `json:"FileUrl,omitempty"`
	CreatedAt   time.Time          `json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `json:"UpdatedAt,omitempty"`
	DeletedAt   time.Time          `json:"DeletedAt,omitempty"`
}
