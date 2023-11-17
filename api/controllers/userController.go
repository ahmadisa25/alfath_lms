package controllers

import (
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	UserController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		userService     interfaces.UserServiceInterface
	}
)

func (userController *UserController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	userService interfaces.UserServiceInterface,
) {
	userController.responder = responder
	userController.customValidator = customValidator
	userController.userService = userService
}

func (userController *UserController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return userController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	user := &models.User{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Password:    funcs.ValidateStringFormKeys("Password", form, "string").(string),
		Email:       funcs.ValidateStringFormKeys("Email", form, "string").(string),
		MobilePhone: funcs.ValidateStringFormKeys("MobilePhone", form, "string").(string),
		CreatedAt:   time.Now(),
	}

	validateError := userController.customValidator.Validate.Struct(user)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(userController.customValidator.TranslateError(validateError))
		fmt.Println(errorResponse)
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return userController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return userController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	if user.Password != "" {
		if len(user.Password) < 8 {
			return userController.responder.HTTP(400, strings.NewReader("Passwords must have a minimum length of 8 characters!"))
		} else {
			user.Password = funcs.HashStringToSHA256(user.Password)
		}
	}

	result, err := userController.userService.Create(*user)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return userController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return userController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return userController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return userController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}
