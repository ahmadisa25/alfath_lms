package middleware

import (
	"alfath_lms/api/funcs"
	"context"
	"os"
	"strings"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	Responder *web.Responder
}

func (authMdw *AuthMiddleware) AuthCheck(ctx context.Context, req *web.Request, action web.Action, prefferedRole []string) web.Result {
	if req.Request().Method == "OPTIONS" {
		//w := new http.ResponseWriter()
		return action(ctx, req)
	}
	headers := req.Request().Header
	//get path = login

	auth, authOk := headers["Authorization"]
	if !authOk {
		return funcs.CorsedResponse(authMdw.Responder.HTTP(401, strings.NewReader("Not authorized")))
	} else {
		auth := strings.Split(auth[0], " ")
		token, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return "Not authorized!", nil
			}

			return []byte(os.Getenv("JWT_KEY")), nil //Parse function must return a key. remember it's called the "Keyfunc".
		})

		if err != nil {
			return funcs.CorsedResponse(authMdw.Responder.HTTP(401, strings.NewReader("Not authorized")))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			claimRole := claims["role_name"].(string)
			if prefferedRole != nil {
				if !funcs.ArrayExists(claimRole, prefferedRole) {
					return funcs.CorsedResponse(authMdw.Responder.HTTP(401, strings.NewReader("Not authorized")))
				}

			}
			req.Request().Header.Add("id", claims["id"].(string))
			req.Request().Header.Add("email", claims["email"].(string))
			req.Request().Header.Add("role_name", claimRole)
		}

		return action(ctx, req)
	}

}
