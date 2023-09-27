package entity

import (
	"time"

	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	ID          int
	Name        string    `gorm:"primaryKey;" validate:"required"`
	Email       string    `gorm:"uniqueIndex" validate:"required,email"`
	MobilePhone string    `gorm:"unique" validate:"required,numeric"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time
}

// TableName overrides the table name used by User to `profiles`
func (Instructor) TableName() string {
	return "ms_instructor"
}
