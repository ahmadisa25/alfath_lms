package models

import (
	"time"

	"gorm.io/gorm"
)

type Announcement struct {
	gorm.Model
	ID          int       `gorm:"primaryKey" json:"ID,omitempty"`
	Name        string    `gorm:"type:varchar(200)" validate:"required" json:"Name,omitempty"`
	Description string    `gorm:"type:varchar(1024)" validate:"required" json:"Description,omitempty"`
	FileUrl     string    `gorm:"type:varchar(500)" json:"FileUrl,omitempty"`
	CreatedAt   time.Time `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time `json:"UpdatedAt,omitempty"`
	DeletedAt   time.Time `gorm:"default:NULL" json:"DeletedAt,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (Announcement) TableName() string {
	return "ms_announcement"
}
