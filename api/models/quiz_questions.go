package models

import (
	"time"

	"gorm.io/gorm"
)

type QuizQuestion struct {
	gorm.Model
	ID            int         `gorm:"primaryKey" json:"ID,omitempty"`
	Title         string      `validate:"required" json:"Title,omitempty"`
	Type          string      `validate:"required" json:"Type,omitempty"`
	Length        int         `validate:"required,numeric" json:"Length,omitempty"`
	Choices       string      `validate:"required" json:"Choices,omitempty"`
	CreatedAt     time.Time   `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt     time.Time   `json:"UpdatedAt,omitempty"`
	DeletedAt     time.Time   `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	ChapterQuiz   ChapterQuiz `validate:"omitempty" json:"ChapterQuiz,omitempty"`
	ChapterQuizID int         `gorm:"foreignKey:ChapterQuizID" validate:"required" json:"ChapterQuizID,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (QuizQuestion) TableName() string {
	return "ms_quiz_questions"
}
