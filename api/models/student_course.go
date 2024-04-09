package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentCourse struct {
	gorm.Model
	ID        int       `gorm:"primaryKey" json:"ID,omitempty"`
	StudentID int       `gorm:"foreignKey:StudentID" validate:"required" json:"StudentID,omitempty"`
	CourseID  int       `gorm:"foreignKey:CourseID" validate:"required" json:"QuizID,omitempty"`
	CreatedAt time.Time `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`
	DeletedAt time.Time `gorm:"default:NULL" json:"DeletedAt,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (StudentCourse) TableName() string {
	return "ms_course_student"
}
