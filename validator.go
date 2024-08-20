package openapi

import (
	"encoding/json"
	"errors"
	openapivalidator "github.com/corytech/go-openapi/validator"
	"github.com/corytech/go-tracing"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var (
	validate            *validator.Validate
	validatorTranslator *ut.Translator
)

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()

		_ = entranslations.RegisterDefaultTranslations(validate, *GetValidatorTranslator())
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})
		openapivalidator.RegisterRFC3339Date(validate, GetValidatorTranslator())
	}
	return validate
}

func GetValidatorTranslator() *ut.Translator {
	if validatorTranslator == nil {
		enTranslator := en.New()
		uni := ut.New(enTranslator, enTranslator)
		trans, _ := uni.GetTranslator("en")
		validatorTranslator = &trans
	}

	return validatorTranslator
}

func BindAndValidate(r any, c *gin.Context) bool {
	if err := c.ShouldBindJSON(r); err != nil {
		var ute *json.UnmarshalTypeError
		if errors.As(err, &ute) {
			NewFromUnmarshalTypeError(ute).Send(c)
			return false
		}

		NewError(BadRequestErrCode).Send(c)
		return false
	}

	if err := GetValidator().Struct(r); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			NewValidationError(ve, GetValidatorTranslator()).Send(c)
			return false
		}
		tracing.HandleErrorForGin(c, err)
		NewError(UnknownErrCode).Send(c)
		return false
	}

	return true
}
