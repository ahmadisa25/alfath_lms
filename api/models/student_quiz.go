package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentQuiz struct {
	gorm.Model
	ID            int         `gorm:"primaryKey" json:"ID,omitempty"`
	Student       Student     `validate:"omitempty" json:"Student,omitempty"`
	StudentID     int         `gorm:"foreignKey:StudentID" validate:"required" json:"StudentID,omitempty"`
	ChapterQuiz   ChapterQuiz `validate:"omitempty" json:"ChapterQuiz,omitempty"`
	ChapterQuizID int         `gorm:"foreignKey:ChapterQuizID" validate:"required" json:"ChapterQuizID,omitempty"`
	CreatedAt     time.Time   `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt     time.Time   `json:"UpdatedAt,omitempty"`
	DeletedAt     time.Time   `gorm:"default:NULL" json:"DeletedAt,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (StudentQuiz) TableName() string {
	return "ms_student_quiz"
}
