package validator

import (
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var v = validator.New()
var english = en.New()
var uni = ut.New(english, english)
var trans, _ = uni.GetTranslator("en")

func init() {
	en_translations.RegisterDefaultTranslations(v, trans)
}

func Validate(s any) error {
	err := v.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Errorf("Invalid validation error: %s", err)
			return errors.Internal()
		}

		for _, err := range err.(validator.ValidationErrors) {
			return errors.InvalidArgument(err.Translate(trans))
		}
	}

	return nil
}
