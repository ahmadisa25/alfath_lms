package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `gorm:"type:varchar(200)" validate:"required" json:"Name,omitempty"`
	Description string             `gorm:"type:varchar(1024)" validate:"required" json:"Description,omitempty"`
	FileUrl     string             `gorm:"type:varchar(500)" json:"FileUrl,omitempty"`
	CreatedAt   time.Time          `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `json:"UpdatedAt,omitempty"`
	DeletedAt   time.Time          `gorm:"default:NULL" json:"DeletedAt,omitempty"`
}
