package cvalid

import (
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chinese "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	Validator      *validator.Validate
	validatorTrans ut.Translator
)

func SetupValid() {
	Validator = binding.Validator.Engine().(*validator.Validate)
	validatorTransHandler()
	registerCustomFieldName()
	registerAll()
}

func registerAll() {
	registerLocationValidatorTrans()
	_ = Validator.RegisterValidation(locationValidatorTag, locationValidator)
}

func validatorTransHandler() {
	zhCn := zh.New()
	uni := ut.New(zhCn, zhCn)
	validatorTrans, _ = uni.GetTranslator("zh")
	_ = chinese.RegisterDefaultTranslations(Validator, validatorTrans)
}

func registerCustomFieldName() {
	Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			return field.Name
		}
		return label
	})
}

type RequestBindingError map[string]string

func FormatError(err error) *RequestBindingError {
	if err, ok := err.(validator.ValidationErrors); ok {
		errors := make(RequestBindingError)
		for _, value := range err {
			errors[value.Field()] = value.Translate(validatorTrans)
		}
		return &errors
	}
	return &RequestBindingError{
		"errors": "字段验证错误",
	}
}

func (req *RequestBindingError) All() map[string]string {
	return *req
}

func (req *RequestBindingError) First() string {
	errorValues := maputil.Values[string, string](*req)
	if len(errorValues) < 1 {
		return ""
	}
	return errorValues[0]
}
