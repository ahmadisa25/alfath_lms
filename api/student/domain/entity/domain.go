package entity

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID          int       `gorm:"primaryKey" json:"ID,omitempty"`
	Name        string    `gorm:"primaryKey;" validate:"required" json:"Name,omitempty"`
	Email       string    `gorm:"unique" validate:"required,email" json:"Email,omitempty"`
	MobilePhone string    `gorm:"unique" validate:"required,numeric" json:"MobilePhone,omitempty"`
	CreatedAt   time.Time `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time `json:"UpdatedAt,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (Student) TableName() string {
	return "ms_student"
}
