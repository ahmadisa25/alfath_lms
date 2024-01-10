package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	InstructorController struct {
		responder         *web.Responder
		customValidator   *validator.CustomValidator
		instructorService interfaces.InstructorServiceInterface
		userService       interfaces.UserServiceInterface
	}

	GetInstructorResponse struct {
		Status int
		Data   models.Instructor
	}
)

func (instructorController *InstructorController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	instructorService interfaces.InstructorServiceInterface,
	userService interfaces.UserServiceInterface,
) {
	instructorController.responder = responder
	instructorController.customValidator = customValidator
	instructorController.instructorService = instructorService
	instructorController.userService = userService
}

func (instructorController *InstructorController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	instructor := &models.Instructor{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Email:       funcs.ValidateStringFormKeys("Email", form, "string").(string),
		MobilePhone: funcs.ValidateStringFormKeys("MobilePhone", form, "string").(string),
		CreatedAt:   time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := instructorController.customValidator.Validate.Struct(instructor)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(instructorController.customValidator.TranslateError(validateError))
		fmt.Println(errorResponse)
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := instructorController.instructorService.CreateInstructor(*instructor)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (instructorController *InstructorController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	instructor, err := instructorController.instructorService.GetInstructor(intID)
	if err != nil {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if instructor.ID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Instructor Not Found!",
		})
	}

	return instructorController.responder.Data(GetInstructorResponse{
		Status: 200,
		Data:   instructor,
	})
}

func (instructorController *InstructorController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	instructor, err := instructorController.instructorService.GetInstructor(intID)
	if err != nil {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if instructor.ID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Instructor Not Found!",
		})
	}

	result, err := instructorController.instructorService.DeleteInstructor(intID)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	_, deleteUserErr := instructorController.userService.Delete(instructor.Email)

	if deleteUserErr != nil {
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(deleteUserErr.Error())))
	}

	return funcs.CorsedResponse(instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (instructorController *InstructorController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return instructorController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		})
	}

	instructor, err := instructorController.instructorService.GetInstructor(intID)
	if err != nil {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if instructor.ID <= 0 {
		return instructorController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Instructor Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	instructorData := &models.Instructor{
		Name:        funcs.ValidateOrOverwriteStringFormKeys("Name", form, "string", instructor).(string),
		Email:       funcs.ValidateOrOverwriteStringFormKeys("Email", form, "string", instructor).(string),
		MobilePhone: funcs.ValidateOrOverwriteStringFormKeys("MobilePhone", form, "string", instructor).(string),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := instructorController.customValidator.Validate.Struct(instructorData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(instructorController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return instructorController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return instructorController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := instructorController.instructorService.UpdateInstructor(intID, *instructorData, instructor)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return instructorController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return instructorController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	instructorController.userService.Update(instructorData.Email, []bson.E{{"email", instructorData.Email}, {"name", instructorData.Name}, {"mobilephone", instructorData.MobilePhone}})

	return instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (instructorController *InstructorController) GetAll(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := instructorController.instructorService.GetAllInstructors(paginationReq)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return instructorController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return funcs.CorsedResponse(instructorController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(instructorController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}
