package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"alfath_lms/api/models"
	"context"
	"encoding/json"
	"strings"
	"time"

	"flamingo.me/flamingo/v3/framework/web"
	"go.mongodb.org/mongo-driver/bson"
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
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	file, handler, err := req.Request().FormFile("file")

	if err != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(err.Error())))
	}

	defer file.Close()

	fileDestination := ""
	if file != nil {
		if funcs.UploadFile(handler.Filename, file) {
			fileDestination = handler.Filename
		}
	}

	announcement := &models.Announcement{
		Title:       funcs.ValidateStringFormKeys("Title", form, "string").(string),
		Description: funcs.ValidateStringFormKeys("Description", form, "string").(string),
		FileUrl:     fileDestination,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	validateError := announcementController.customValidator.Validate.Struct(announcement)
	if validateError != nil {
		errorResponse := funcs.ErrorPackagingForMaps(announcementController.customValidator.TranslateError(validateError))
		errorResponse, packError := funcs.ErrorPackaging(errorResponse, 400)
		if packError != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(errorResponse)))
	}

	result, err := announcementController.announcementService.Create(*announcement)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(announcementController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))

}

func (announcementController *AnnouncementController) GetAll(ctx context.Context, req *web.Request) web.Result {
	query := req.QueryAll()
	paginationReq := definitions.PaginationRequest{
		SelectedColumns: funcs.ValidateStringFormKeys("select", query, "string").(string),
		Search:          funcs.ValidateStringFormKeys("search", query, "string").(string),
		Page:            funcs.ValidateStringFormKeys("page", query, "string").(string),
		PerPage:         funcs.ValidateStringFormKeys("perpage", query, "string").(string),
		OrderBy:         funcs.ValidateStringFormKeys("order", query, "string").(string),
		Filter:          funcs.ValidateStringFormKeys("filter", query, "string").(string),
	}

	result, err := announcementController.announcementService.GetAll(paginationReq)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(announcementController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (announcementController *AnnouncementController) Get(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(announcementController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an instructor!",
		}))
	}

	announcement, err := announcementController.announcementService.Get(req.Params["id"])
	if err != nil {
		return funcs.CorsedDataResponse(announcementController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	return funcs.CorsedDataResponse(announcementController.responder.Data(announcement))
}

func (announcementController *AnnouncementController) Delete(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(announcementController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an announcement!",
		}))
	}

	result, err := announcementController.announcementService.Delete(req.Params["id"])
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(announcementController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}

func (announcementController *AnnouncementController) Update(ctx context.Context, req *web.Request) web.Result {
	if req.Params["id"] == "" {
		return funcs.CorsedDataResponse(announcementController.responder.Data(definitions.GenericAPIMessage{
			Status:  400,
			Message: "Please select an announcement!",
		}))
	}

	formError := req.Request().ParseForm()
	if formError != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(formError.Error())))
	}

	form := req.Request().Form

	existingAnnouncement, err := announcementController.announcementService.Get(req.Params["id"])
	if err != nil {
		return funcs.CorsedDataResponse(announcementController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	fileDestination := existingAnnouncement.Data.FileUrl

	file, handler, err := req.Request().FormFile("file")

	if file != nil {
		if err != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(err.Error())))
		}

		defer file.Close()

		if file != nil {
			if funcs.UploadFile(handler.Filename, file) {
				fileDestination = handler.Filename
			}
		}
	}

	announcement := []bson.E{
		{"title", funcs.ValidateOrOverwriteStringFormKeys("Title", form, "string", existingAnnouncement.Data).(string)},
		{"description", funcs.ValidateOrOverwriteStringFormKeys("Description", form, "string", existingAnnouncement.Data).(string)},
		{"fileurl", fileDestination},
		{"updatedat", time.Now()},
	}

	result, err := announcementController.announcementService.Update(req.Params["id"], announcement)
	if err != nil {
		errorResponse, packError := funcs.ErrorPackaging(err.Error(), 500)
		if packError != nil {
			return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(packError.Error())))
		}
		return funcs.CorsedResponse(announcementController.responder.HTTP(500, strings.NewReader(errorResponse)))
	}

	res, resErr := json.Marshal(result)
	if resErr != nil {
		return funcs.CorsedResponse(announcementController.responder.HTTP(400, strings.NewReader(resErr.Error())))
	}

	return funcs.CorsedResponse(announcementController.responder.HTTP(uint(result.Status), strings.NewReader(string(res))))
}
