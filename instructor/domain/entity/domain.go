package entity

import (
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	ID          int
	Name        string	`validate:"required"`
	Email       string	`validate:"required,email"`
	MobilePhone string	`validate:"required"`
}
  
// TableName overrides the table name used by User to `profiles`
func (Instructor) TableName() string {
	return "ms_instructor"
}
