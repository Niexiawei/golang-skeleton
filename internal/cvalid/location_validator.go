package cvalid

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

const locationValidatorTag = "location"

var (
	lonRegex = regexp.MustCompile("^(\\-|\\+)?(((\\d|[1-9]\\d|1[0-7]\\d|0{1,3})\\.\\d{0,6})|(\\d|[1-9]\\d|1[0-7]\\d|0{1,3})|180\\.0{0,6}|180)$")
	latRegex = regexp.MustCompile("^(\\-|\\+)?([0-8]?\\d{1}\\.\\d{0,6}|90\\.0{0,6}|[0-8]?\\d{1}|90)$")
)

func locationValidator(fl validator.FieldLevel) bool {
	local, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	localArr := strings.Split(local, ",")
	if len(localArr) < 2 {
		return false
	}
	return lonRegex.MatchString(localArr[0]) && latRegex.MatchString(localArr[1])
}

func registerLocationValidatorTrans() {
	_ = Validator.RegisterTranslation(locationValidatorTag, validatorTrans, func(ut ut.Translator) error {
		return ut.Add("name", "{0}坐标格式不正确", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("name", fe.Field())
		return t
	})
}
