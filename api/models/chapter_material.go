package models

import (
	"time"

	"gorm.io/gorm"
)

type ChapterMaterial struct {
	gorm.Model
	ID              int           `gorm:"primaryKey" json:"ID,omitempty"`
	Name            string        `gorm:"type:varchar(200)" validate:"required" json:"Name,omitempty"`
	Description     string        `gorm:"type:varchar(1024)" validate:"required" json:"Description,omitempty"`
	FileUrl         string        `gorm:"type:varchar(500)" validate:"required" json:"FileUrl,omitempty"`
	CreatedAt       time.Time     `gorm:"default:NULL" json:"CreatedAt,omitempty"`
	UpdatedAt       time.Time     `json:"UpdatedAt,omitempty"`
	DeletedAt       time.Time     `gorm:"default:NULL" json:"DeletedAt,omitempty"`
	CourseChapter   CourseChapter `validate:"omitempty" json:"CourseChapter,omitempty"`
	CourseChapterID int           `gorm:"foreignKey:CourseChapterID" validate:"required" json:"CourseChapterID,omitempty"`
}

// TableName overrides the table name used by User to `profiles`
func (ChapterMaterial) TableName() string {
	return "ms_chapter_material"
}
