package validator

import (
	"fmt"

	"flamingo.me/dingo"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Module struct{}

type CustomValidator struct {
	Validate *validator.Validate //lower case field names = protected fields, i.e can only be accessed in this package.
	Trans    *ut.Translator
}

// follwoign the steps from Renaldy @ Tunaiku tech
func (cv *CustomValidator) TranslateError(err error) (errs []error) {
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)

	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(*cv.Trans))
		errs = append(errs, translatedErr)
	}

	return errs
}

func (*Module) Configure(injector *dingo.Injector) {
	injector.Bind(new(CustomValidator)).ToProvider(func() *CustomValidator {
		validate := validator.New()

		//following steps from Renaldi @ Tunaiku Tech
		english := en.New()
		uni := ut.New(english, english)
		trans, _ := uni.GetTranslator("en")
		_ = en_translations.RegisterDefaultTranslations(validate, trans)

		return &CustomValidator{
			Validate: validate,
			Trans:    &trans,
		}
	}).In(dingo.Singleton) //SIngleton = make sure the bind stuff can be used globally.
}