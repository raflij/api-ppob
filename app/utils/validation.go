package utils

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ResultValidation struct {
	Errors ResponseValidation `json:"errors"`
}
type ResponseValidation struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

func ValidationCustom(c *gin.Context, dataInput interface{}) ([]ResponseValidation, bool) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	checkValidate := validate.Struct(dataInput)
	if checkValidate != nil {
		return ResResult(checkValidate), true
	}

	return nil, false
}

func ResResult(checkValidate interface{}) []ResponseValidation {
	validationError := checkValidate.(validator.ValidationErrors)
	valError := make([]ResponseValidation, len(validationError))
	for i, err := range validationError {
		valError[i] = ResponseValidation{
			Message: GetTag(err),
			Field:   err.Field(),
		}
	}
	return valError
}

func GetTag(fe validator.FieldError) string {
	data := map[string]string{
		"required": "harus diisi",
		"min":      "minimal terdiri dari " + fe.Param() + " karakter",
		"max":      "maximal terdiri dari " + fe.Param() + " karakter",
		"number":   "hanya boleh terdiri dari angka",
		"email":    "email tidak valid",
		"alpha":    "hanya boleh terdiri dari huruf",
		"alphanum": "hanya boleh terdiri dari huruf dan angka",
		"eqfield":  "tidak sama dengan " + fe.Param(),
		"unique":   "sudah ada",
	}
	// return nil
	return data[fe.Tag()]
}
