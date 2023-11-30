package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"
)

type UserServiceInterface interface {
	Create(User models.User, Role string) (definitions.GenericMongoCreationMessage, error)
	Delete(Email string) (definitions.GenericAPIMessage, error)
	Login(Data map[string]interface{}) (definitions.LoginResponse, error)
	//Refresh(Data map[string]interface{}) (definitions.LoginResponse, error)
}
