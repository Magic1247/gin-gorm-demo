package utils

import "github.com/go-playground/validator/v10"

func StartWith(fl validator.FieldLevel) bool {
	value := fl.Field().String() // 获取字段的实际值
	param := fl.Param()          // 获取传递的参数

	// 判断字段值是否以参数开头
	return len(value) >= len(param) && value[:len(param)] == param
}
