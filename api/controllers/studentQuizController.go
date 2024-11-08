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
	StudentQuizController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		stdQuizService  interfaces.StudentQuizServiceInterface
	}
)

func (stdQuizController *StudentQuizController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	studentQuizService interfaces.StudentQuizServiceInterface,
) {
	stdQuizController.responder = responder
	stdQuizController.customValidator = customValidator
	stdQuizController.stdQuizService = studentQuizService
}

func (stdQuizController *StudentQuizController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["student_id"] == "" {
		return funcs.CorsedDataResponse(stdQuizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a student!",
		}))
	}

	intID, err := strconv.Atoi(req.Params["student_id"])
	//PrintError(err)

	if intID <= 0 {
		return funcs.CorsedDataResponse(stdQuizController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a student!",
		}))
	}

	answer, err := stdQuizController.stdQuizService.Get(intID)
	if err != nil {
		return funcs.CorsedDataResponse(stdQuizController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	if answer.ID <= 0 {
		return funcs.CorsedDataResponse(stdQuizController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Quiz Data Not Found!",
		}))
	}

	return funcs.CorsedDataResponse(stdQuizController.responder.Data(definitions.GenericGetMessage[models.StudentQuiz]{
		Status: 200,
		Data:   answer,
	}))
}

func (stdQuizController *StudentQuizController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(stdQuizController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	GradedByID := funcs.ValidateStringFormKeys("GradedByID", form, "string").(string)
	var GradedAt time.Time
	if GradedByID != "" {
		GradedAt = time.Now()
	} else {
		GradedAt = time.Time{}
	}

	stdQuiz := &models.StudentQuiz{

		StudentID:  funcs.ValidateStringFormKeys("StudentID", form, "int").(int),
		QuizID:     funcs.ValidateStringFormKeys("QuizID", form, "int").(int),
		FinalGrade: funcs.ValidateStringFormKeys("FinalGrade", form, "int").(int),
		GradedByID: GradedByID,
		GradedAt:   GradedAt,
		CreatedAt:  time.Now(),
	}

	//fmt.Printf("validator: %+v\n", instructorController.validator.validate)
	validateError := stdQuizController.customValidator.Validate.Struct(stdQuiz)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(stdQuizController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(stdQuizController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(stdQuizController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := stdQuizController.stdQuizService.Create(*stdQuiz)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(stdQuizController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(stdQuizController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(stdQuizController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(stdQuizController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}
