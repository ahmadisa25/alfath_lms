package models

import (
	"time"

	"gorm.io/gorm"
)

type QuizAnswer struct {
	gorm.Model
	ID             int          `gorm:"primaryKey" json:"ID,omitempty"`
	Answer         string       `validate:"required" json:"Answer,omitempty"`
	Grade          int          `json:"Grade,omitempty"`
	Student        Student      `validate:"omitempty" json:"Student,omitempty"`
	StudentID      int          `gorm:"foreignKey:StudentID" validate:"required" json:"StudentID,omitempty"`
	CreatedAt      time.Time    `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt      time.Time    `json:"UpdatedAt,omitempty"`
	DeletedAt      time.Time    `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	QuizQuestion   QuizQuestion `validate:"omitempty" json:"QuizQuestion,omitempty"`
	QuizQuestionID int          `gorm:"foreignKey:QuizQuestionID" validate:"required" json:"QuizQuestionID,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (QuizAnswer) TableName() string {
	return "ms_quiz_answers"
}
