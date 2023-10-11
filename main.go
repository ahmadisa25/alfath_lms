package main

import (
	"alfath_lms/api"
	"fmt"

	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	flamingo.App([]dingo.Module{
		new(requestlogger.Module),
		new(api.Module),
	},
	)
}
