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
	AnnouncementController struct {
		responder           *web.Responder
		customValidator     *validator.CustomValidator
		announcementService interfaces.AnnouncementServiceInterface
	}
)

func (announcementController *AnnouncementController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	announcementService interfaces.AnnouncementServiceInterface,
) {
	announcementController.responder = responder
	announcementController.customValidator = customValidator
	announcementController.announcementService = announcementService
}

func (announcementController *AnnouncementController) Create(ctx context.Context, req *web.Request) web.Result {
	formError := req.Request().ParseForm()
	if formError != nil {
		return announcementController.responder.HTTP(400, strings.NewReader(formError.Error()))
	}

	form := req.Request().Form

	announcement := &models.Announcement{
		Title:       funcs.ValidateStringFormKeys("Title", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		FileUrl:     funcs.ValidateStringFormKeys("FileUrl", form, "string").(string),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	validateError := announcementController.customValidator.Validate.Struct(announcement)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(announcementController.customValidator.TranslateError(validateError))
		fmt.Println(errorResponse)
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return announcementController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return announcementController.responder.HTTP(400, strings.NewReader(errorResponse))
	}

	result, err := announcementController.announcementService.Create(*announcement)
	if err != nil {
		fmt.Println(err)
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return announcementController.responder.HTTP(500, strings.NewReader(packError.Error()))
		}
		return announcementController.responder.HTTP(500, strings.NewReader(errorResponse))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return announcementController.responder.HTTP(400, strings.NewReader(resErr.Error()))
	}

	return announcementController.responder.HTTP(uint(result.Status), strings.NewReader(string(res)))

}
