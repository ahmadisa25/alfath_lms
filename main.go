package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/requestlogger"
)

func main() {
	flamingo.App(
		[]dingo.module{
			new(requestlogger.Module),
		},
	)
}
