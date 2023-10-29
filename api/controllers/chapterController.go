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
)

type (
	ChapterController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		chapterService  interfaces.ChapterServiceInterface
	}

	GetChapterResponse struct {
		Status int
		Data   models.CourseChapter
	}
)

func (chapterController *ChapterController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	chapterService interfaces.ChapterServiceInterface,
) {
	chapterController.responder = responder
	chapterController.customValidator = customValidator
	chapterController.chapterService = chapterService
}

func (chapterController *ChapterController) Create(ctx context.Context, req *web.Request) web.Result {
	fmt.Println("test")
	formError := req.Request().ParseForm()
	if formError != nil {
		return chapterController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	chapter := &models.CourseChapter{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:    funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		CourseID:    funcs.ValidateStringFormKeys("CourseID", form, "int").(int),
		CreatedAt:   time.Now(),
	}

	//fmt.Printf("validator: %+v\n", chapterController.validator.validate)
	validateError := chapterController.customValidator.Validate.Struct(chapter)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(chapterController.customValidator.TranslateError(validateError))
		fmt.Println(errorResponse)
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return chapterController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return chapterController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := chapterController.chapterService.Create(*chapter)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return chapterController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return chapterController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return chapterController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return chapterController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}

func (chapterController *ChapterController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an chapter!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select a chapter!",
		})
	}

	chapter, err := chapterController.chapterService.Get(intID)
	if err != nil {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if chapter.ID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "Chapter Not Found!",
		})
	}

	return chapterController.responder.Data(GetChapterResponse{
		Status: 200,
		Data:   chapter,
	})
}

func (chapterController *ChapterController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an chapter!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	//PrintError(err)

	if intID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an chapter!",
		})
	}

	chapter, err := chapterController.chapterService.Get(intID)
	if err != nil {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if chapter.ID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "chapter Not Found!",
		})
	}

	result, err := chapterController.chapterService.Delete(intID)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return chapterController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return chapterController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return chapterController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return chapterController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}

func (chapterController *ChapterController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an chapter!",
		})
	}

	intID, err := strconv.Atoi(req.Params["id"])
	if err != nil {
		return chapterController.responder.HTTP(500, strings.NewReader(err.Error()))
	}
	//PrintError(err)

	if intID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an chapter!",
		})
	}

	chapter, err := chapterController.chapterService.Get(intID)
	if err != nil {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		})
	}

	if chapter.ID <= 0 {
		return chapterController.responder.Data(definitions.GenericAPIMessage{
			Status:  404,
			Message: "chapter Not Found!",
		})
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return chapterController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	chapterData := &models.CourseChapter{
		Name:        funcs.ValidateStringFormKeys("Name", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		Duration:    funcs.ValidateStringFormKeys("Duration", form, "int").(int),
		CourseID:    funcs.ValidateStringFormKeys("CourseID", form, "int").(int),
		CreatedAt:   time.Now(),
	}

	//fmt.Printf("validator: %+v\n", chapterController.validator.validate)
	validateError := chapterController.customValidator.Validate.Struct(chapterData)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(chapterController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return chapterController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return chapterController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := chapterController.chapterService.Update(intID, *chapterData)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return chapterController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return chapterController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return chapterController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return chapterController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))
}
