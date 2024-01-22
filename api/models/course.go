package models

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID          int           `gorm:"primaryKey" json:"ID,omitempty"`
	Name        string        `validate:"required" json:"Name,omitempty"`
	Description string        `validate:"required" json:"Description,omitempty"`
	Duration    int           `validate:"required,numeric" json:"Duration,omitempty"`
	FileUrl     string        `json:"FileUrl,omitempty"`
	CreatedAt   time.Time     `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time     `json:"UpdatedAt,omitempty"`
	DeletedAt   time.Time     `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	Instructors []*Instructor `gorm:"many2many:ms_course_instructor" json:"Instructors,omitempty"`
	Students    []*Student    `gorm:"many2many:ms_course_student" json:"Students,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (Course) TableName() string {
	return "ms_course"
}
