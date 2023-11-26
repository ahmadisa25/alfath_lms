package middleware

import (
	"context"
	"strings"

	"flamingo.me/flamingo/v3/framework/web"
)

type AuthMiddleware struct {
	Responder *web.Responder
}

func (authMdw *AuthMiddleware) AuthCheck(ctx context.Context, req *web.Request, action web.Action) web.Result {
	headers := req.Request().Header
	//get path = login

	if headers["Authorization"] == nil {
		return authMdw.Responder.HTTP(401, strings.NewReader("Not authorized!"))
	} else {
		return action(ctx, req)
	}

}
