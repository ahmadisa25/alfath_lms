package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type UserServiceInterface interface {
	Create(User models.User) (definitions.GenericMongoCreationMessage, error)
}
