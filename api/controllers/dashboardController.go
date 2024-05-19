package controllers

import (
	"alfath_lms/api/definitions"
	"alfath_lms/api/deps/validator"
	"alfath_lms/api/funcs"
	"alfath_lms/api/interfaces"
	"context"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	DashboardController struct {
		responder       *web.Responder
		customValidator *validator.CustomValidator
		dashService     interfaces.DashboardServiceInterface
	}
)

func (dashController *DashboardController) Inject(
	responder *web.Responder,
	customValidator *validator.CustomValidator,
	dashService interfaces.DashboardServiceInterface,
) {
	dashController.responder = responder
	dashController.customValidator = customValidator
	dashController.dashService = dashService
}

func (dashController *DashboardController) GetDashboardData(ctx context.Context, req *web.Request) web.Result {
	data, err := dashController.dashService.GetDashboardData()
	if err != nil {
		return funcs.CorsedDataResponse(dashController.responder.Data(definitions.GenericAPIMessage{
			Status:  500,
			Message: "We cannot process your request. Please try again or contact support!",
		}))
	}

	return funcs.CorsedDataResponse(dashController.responder.Data(definitions.GenericGetMessage[definitions.SimpleDashboardData]{
		Status: 200,
		Data:   data,
	}))
}
