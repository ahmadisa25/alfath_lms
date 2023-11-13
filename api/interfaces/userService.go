package interfaces

import (
	"alfath_lms/api/definitions"
)

type UserServiceInterface interface {
	Create(data map[string]interface{}) (definitions.GenericCreationMessage, error)
}
