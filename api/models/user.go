package models

import (
	"time"
)

type User struct {
	ID          string    `json:"ID,omitempty"`
	Name        string    `validate:"required" json:"Name,omitempty"`
	Password    string    `validate:"required" json:"Password,omitempty"`
	Email       string    `validate:"required,email" json:"Email,omitempty"`
	MobilePhone string    `validate:"required,numeric" json:"MobilePhone,omitempty"`
	CreatedAt   time.Time `json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time `json:"UpdatedAt,omitempty"`
}
