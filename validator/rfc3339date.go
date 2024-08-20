package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"time"
)

func RegisterRFC3339Date(v *validator.Validate, t *ut.Translator) {
	_ = v.RegisterValidation("IsRFC3339Date", isRFC3339Date)
	_ = v.RegisterTranslation(
		"IsRFC3339Date",
		*t,
		func(trans ut.Translator) error {
			if err := trans.Add("IsRFC3339Date", "{0} must be in RFC3339 datetime format", false); err != nil {
				return err
			}

			return nil
		},
		func(trans ut.Translator, fe validator.FieldError) string {
			msg, err := trans.T(fe.Tag(), fe.Field())
			if err != nil {
				return ""
			}

			return msg
		})
}

func isRFC3339Date(fl validator.FieldLevel) bool {
	v := fl.Field().String()

	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return false
	}

	return true
}
