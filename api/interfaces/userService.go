package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/models"

	"go.mongodb.org/mongo-driver/bson"
)

type UserServiceInterface interface {
	Create(User models.User, Role string) (definitions.GenericMongoCreationMessage, error)
	Update(Email string, Updates []bson.E) (definitions.GenericAPIMessage, error)
	Delete(Email string) (definitions.GenericAPIMessage, error)
	Login(Data map[string]interface{}) (definitions.LoginResponse, error)
	LoginAdmin(Data map[string]interface{}) (definitions.LoginResponse, error)
	Refresh(Data map[string]interface{}) (definitions.LoginResponse, error)
}
