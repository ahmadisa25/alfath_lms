package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentQuiz struct {
	gorm.Model
	ID int `gorm:"primaryKey" json:"ID,omitempty"`
	//Student    Student     `json:"Student,omitempty"`
	StudentID int `gorm:"foreignKey:StudentID" validate:"required" json:"StudentID,omitempty"`
	//Quiz       ChapterQuiz `json:"ChapterQuiz,omitempty"`
	QuizID     int `gorm:"foreignKey:QuizID" validate:"required" json:"QuizID,omitempty"`
	FinalGrade int `json:"FinalGrade,omitempty" validate:"required"`
	//GradedBy   Instructor ` json:"GradedBy,omitempty"`
	GradedByID int       `validate:"required" json:"GradedByID,omitempty"` //Instructor
	GradedAt   time.Time `gorm:"default:NULL" json:"GradedAt,omitempty" validate:"required"`
	CreatedAt  time.Time `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt  time.Time `json:"UpdatedAt,omitempty"`
	DeletedAt  time.Time `gorm:"default:NULL" json:"DeletedAt,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (StudentQuiz) TableName() string {
	return "ms_student_quiz"
}
