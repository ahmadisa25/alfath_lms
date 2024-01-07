package controllers

import (
	"context"
	"strings"

	"flamingo.me/flamingo/v3/framework/web"
)

type (
	OptionsHandler struct {
		responder *web.Responder
	}
)

func (optionsHandler *OptionsHandler) Inject(
	responder *web.Responder,
) {
	optionsHandler.responder = responder
}

func (optionsHandler *OptionsHandler) Setup(ctx context.Context, req *web.Request) web.Result {
	res := optionsHandler.responder.HTTP(200, strings.NewReader("allowed"))
	res.Header.Add("Access-Control-Allow-Origin", "*")
	res.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	res.Header.Add("Access-Control-Allow-Headers", "*")
	res.Header.Add("Access-Control-Max-Age", "86400")
	return res
}
