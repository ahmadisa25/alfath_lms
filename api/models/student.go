package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	ID          int       `gorm:"primaryKey" json:"ID,omitempty"`
	Name        string    `gorm:"type:varchar(200)" validate:"required" json:"Name,omitempty"`
	Email       string    `gorm:"unique,type:varchar(200)" validate:"required,email" json:"Email,omitempty"`
	MobilePhone string    `gorm:"unique" validate:"required,numeric" json:"MobilePhone,omitempty"`
	CreatedAt   time.Time `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time `json:"UpdatedAt,omitempty"`
	Courses     []*Course `gorm:"many2many:ms_course_student"`
}

// TableName overrides the table name used by User to `profiles`
func (Student) TableName() string {
	return "ms_student"
}
