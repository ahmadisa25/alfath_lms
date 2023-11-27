package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"flamingo.me/flamingo/v3/framework/web"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	Responder *web.Responder
}

func (authMdw *AuthMiddleware) AuthCheck(ctx context.Context, req *web.Request, action web.Action) web.Result {
	headers := req.Request().Header
	//get path = login

	auth, authOk := headers["Authorization"]
	if !authOk {
		return authMdw.Responder.HTTP(401, strings.NewReader("Not authorized!"))
	} else {
		auth := strings.Split(auth[0], " ")
		token, err := jwt.Parse(auth[1], func(token *jwt.Token) (interface{}, error) {
			val, ok := token.Method.(*jwt.SigningMethodHMAC)
			fmt.Println(val)
			if !ok {
				return "Not authorized!", nil
			}

			return []byte(os.Getenv("JWT_KEY")), nil //Parse function must return a key. remember it's called the "Keyfunc".
		})

		if err != nil {
			return authMdw.Responder.HTTP(401, strings.NewReader("Not authorized!"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			req.Request().Header.Add("email", claims["email"].(string))
			req.Request().Header.Add("role_name", claims["role_name"].(string))
		}

		fmt.Println(req.Request().Header)
		return action(ctx, req)
	}

}
