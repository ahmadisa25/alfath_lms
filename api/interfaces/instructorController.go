package interfaces

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/funcs"
	"alfath_lms/deps/validator"
	"alfath_lms/instructor/domain/entity"
	"alfath_lms/instructor/domain/service"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	InstructorController struct {
		responder         *web.Responder
		validator         *validator.CustomValidator
		instructorService service.InstructorServiceInterface
	}

	GetInstructorResponse struct {
		Status int
		Data   entity.Instructor
	}
)

func (instructorController *InstructorController) Inject(
	responder *web.Responder,
	validator *validator.CustomValidator,
	instructorService service.InstructorServiceInterface,
) {
	instructorController.responder = responder
	instructorController.validator = validator
	instructorController.instructorService = instructorService
}

func (instructorController *InstructorController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	instructor := &entity.Instructor{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Email:       funcs.ValidateStringFormKeys("Email", form, "string").(string),
		MobilePhone: funcs.ValidateStringFormKeys("MobilePhone", form, "string").(string),
		CreatedAt:   time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := instructorController.validator.Validate.Struct(instructor)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(instructorController.validator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return instructorController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return instructorController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := instructorController.instructorService.CreateInstructor(*instructor)
	if err != nil {
		fmt.Println(err)
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

	return instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}

func (instructorController *InstructorController) Update(ctx context.Context, req *web.Request) web.Result {
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

	formError := req.Request().ParseForm()
	if formError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	instructorData := &entity.Instructor{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Email:       funcs.ValidateStringFormKeys("Email", form, "string").(string),
		MobilePhone: funcs.ValidateStringFormKeys("MobilePhone", form, "string").(string),
		CreatedAt:   time.Now(),
	}

	validateError := instructorController.validate.Struct(instructorData)
	if validateError != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(validateError.Error()))
	}

	result, err := instructorController.instructorService.CreateInstructor(*instructorData)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return instructorController.responder.HTTP(500, strings.NewReader(err.Error()))
		}
		return instructorController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return instructorController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return instructorController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

	/*return instructorController.responder.Data(GetInstructorResponse{
		Status: 200,
		Data:   instructor,
	})*/
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

func PrintError(err error) error {
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
