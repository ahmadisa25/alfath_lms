package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	StudentCourseController struct {
		responder        *web.Responder
		customValidator  *validator.CustomValidator
		stdCourseService interfaces.StudentCourseServiceInterface
	}
)

func (stdCourseController *StudentCourseController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	studentCourseService interfaces.StudentCourseServiceInterface,
) {
	stdCourseController.responder = responder
	stdCourseController.customValidator = customValidator
	stdCourseController.stdCourseService = studentCourseService
}

func (stdCourseController *StudentCourseController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	stdCourse := &models.StudentCourse{

		StudentID: funcs.ValidateStringFormKeys("StudentID", form, "int").(int),
		CourseID:  funcs.ValidateStringFormKeys("CourseID", form, "int").(int),
		CreatedAt: time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := stdCourseController.customValidator.Validate.Struct(stdCourse)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(stdCourseController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(stdCourseController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := stdCourseController.stdCourseService.Create(*stdCourse)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(stdCourseController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(stdCourseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (stdCourseController *StudentCourseController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(stdCourseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an quiz!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if intID <= 0 {
		return funcs.CorsedDataResponse(stdCourseController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select student!",
		}))
	}

	result, err := stdCourseController.stdCourseService.Delete(intID)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(stdCourseController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(stdCourseController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(stdCourseController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}
