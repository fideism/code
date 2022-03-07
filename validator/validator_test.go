package validator

import (
	"fmt"
	"reflect"
	"testing"
)

// User ....
type User struct {
	IntRange       int
	IntMin         int
	StringLength   string
	StringRequired string
}

// UserValidate ...
type UserValidate struct {
	IntRange       int    `validate:"int,min=1,max=1000"`
	IntMin         int    `validate:"int,min=10"`
	StringLength   string `validate:"string,min=2,max=10"`
	StringRequired string `validate:"string,min=1"`
}

func Test_Validator(t *testing.T) {
	userData := User{
		IntRange:       10,
		IntMin:         0,
		StringRequired: "111",
		StringLength:   "foobar",
	}
	deal(&userData, &UserValidate{})
}

func deal(origin, validate interface{}) {
	originKeyElem := reflect.TypeOf(origin).Elem()
	validateKeyElem := reflect.TypeOf(validate).Elem()

	originElem := reflect.ValueOf(origin).Elem()
	validateElem := reflect.ValueOf(validate).Elem()

	for i := 0; i < originKeyElem.NumField(); i++ {
		for j := 0; j < validateKeyElem.NumField(); j++ {

			if originKeyElem.Field(i).Name == validateKeyElem.Field(j).Name {
				validateElem.Field(j).Set(originElem.Field(j))
			}

		}
	}

	for i, err := range ValidateStruct(validate) {
		fmt.Println(i, err.Error())
	}
}
