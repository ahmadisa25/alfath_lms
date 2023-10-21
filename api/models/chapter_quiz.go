package models

import (
	"time"

	"gorm.io/gorm"
)

type ChapterQuiz struct {
	gorm.Model
	ID              int           `gorm:"primaryKey" json:"ID,omitempty"`
	Name            string        `validate:"required" json:"Name,omitempty"`
	Description     string        `validate:"required" json:"Description,omitempty"`
	Duration        int           `validate:"required,numeric" json:"Duration,omitempty"`
	CreatedAt       time.Time     `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt       time.Time     `json:"UpdatedAt,omitempty"`
	DeletedAt       time.Time     `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	CourseChapter   CourseChapter `validate:"omitempty" json:"CourseChapter,omitempty"`
	CourseChapterID int           `gorm:"foreignKey:CourseChapterID" validate:"required" json:"CourseChapterID,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (ChapterQuiz) TableName() string {
	return "ms_chapter_quiz"
}
