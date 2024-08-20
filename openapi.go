package openapi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type ValidationErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Error struct {
	Code             ErrorCode              `json:"code"`
	ValidationErrors *[]ValidationErrorItem `json:"validationErrors"`
}

type ResponseWrapper struct {
	Error *Error `json:"error"`
	Data  any    `json:"data"`
}

func NewOk(data any) ResponseWrapper {
	return ResponseWrapper{
		Error: nil,
		Data:  data,
	}
}

func NewError(code ErrorCode) ResponseWrapper {
	return ResponseWrapper{
		Error: &Error{
			Code:             code,
			ValidationErrors: nil,
		},
		Data: nil,
	}
}

func NewFromUnmarshalTypeError(ute *json.UnmarshalTypeError) ResponseWrapper {
	return ResponseWrapper{
		Error: &Error{
			Code: ValidationErrCode,
			ValidationErrors: &[]ValidationErrorItem{
				{
					Field:   ute.Field,
					Message: fmt.Sprintf("Cannot unmarshal value. `%s` expected, `%s` given", ute.Type.String(), ute.Value),
				},
			},
		},
		Data: nil,
	}
}

func NewValidationError(ve validator.ValidationErrors, trans *ut.Translator) ResponseWrapper {
	validationErrors := make([]ValidationErrorItem, len(ve))

	for i, e := range ve {
		validationErrors[i] = ValidationErrorItem{
			Field:   strings.Join(strings.Split(e.Namespace(), ".")[1:], "."),
			Message: e.Translate(*trans),
		}
	}
	return ResponseWrapper{
		Error: &Error{
			Code:             ValidationErrCode,
			ValidationErrors: &validationErrors,
		},
		Data: nil,
	}
}

func NotFound(ctx *gin.Context) {
	NewError(MethodNotFoundErrCode).Send(ctx)
}

func MethodNotAllowed(ctx *gin.Context) {
	NewError(InvalidHttpMethodErrCode).Send(ctx)
}

func (r ResponseWrapper) Send(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
}
