package models

import (
	"time"

	"gorm.io/gorm"
)

type CourseChapter struct {
	gorm.Model
	ID          int         `gorm:"primaryKey" json:"ID,omitempty"`
	Name        string      `validate:"required" json:"Name,omitempty"`
	Description string      `validate:"required" json:"Description,omitempty"`
	Duration    int         `validate:"required,numeric" json:"Duration,omitempty"`
	CreatedAt   time.Time   `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time   `json:"UpdatedAt,omitempty"`
	DeletedAt   time.Time   `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	Course      Course      `validate:"omitempty" json:"Course,omitempty"`
	CourseID    int         `gorm:"foreignKey:CourseID" validate:"required" json:"CourseID,omitempty"`
	Materials   []Materials `json:"Materials,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (CourseChapter) TableName() string {
	return "ms_course_chapters"
}
