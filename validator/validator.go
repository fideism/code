package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// Validator 验证接口
type Validator interface {
	Validate(interface{}) (bool, error)
}

// DefaultValidator 默认
type DefaultValidator struct{}

// Validate 默认
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

const tagName = "validate"

// ValidateStruct 校验结构体
func ValidateStruct(s interface{}) []error {
	var errs []error
	v := reflect.ValueOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)

		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "int":
		validator := IntValidator{}
		//将structTag中的min和max解析到结构体中
		_, _ = fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		_, _ = fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	}

	return DefaultValidator{}
}

// StringValidator 字符串长度校验规则...
type StringValidator struct {
	Min int
	Max int
}

// Validate ...
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("字符串不能为空")
	}

	if l < v.Min {
		return false, fmt.Errorf("字符长度必须大于 %v", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("字符长度必须小于 %v", v.Max)
	}

	return true, nil
}

// IntValidator 数字校验
type IntValidator struct {
	Min int
	Max int
}

// Validate ...
func (v IntValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)

	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}
