package db

import (
	"flamingo.me/dingo"
	"github.com/go-playground/validator/v10"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector){	
	injector.Bind(validator.New()).ToProvider( func() *validator.Validate {
		validate := validator.New()
		return validate
	}).In(dingo.Singleton) //SIngleton = make sure the bind stuff can be used globally.
}