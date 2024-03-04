package api

import (
	"github.com/Yadier01/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCurency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if curreny, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(curreny)
	}
	return false
}
