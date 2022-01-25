package binding

// 抄 https://github.com/gin-gonic/gin/blob/master/binding/binding.go

import (
	"encoding/json"
	"io"
	"net/url"
	"reflect"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	// valuerDecoder 声明为包级变量
	valuerDecoder = schema.NewDecoder()
)

func init() {
	valuerDecoder.IgnoreUnknownKeys(true)
	valuerDecoder.RegisterConverter(uuid.NewV4(), schema.Converter(func(value string) reflect.Value {
		if v, err := uuid.FromString(value); nil == err {
			return reflect.ValueOf(v)
		}
		return reflect.Value{}
	}))

	govalidator.CustomTypeTagMap.Set("slice.string", govalidator.CustomTypeValidator(SliceString))
	govalidator.CustomTypeTagMap.Set("notnull", govalidator.CustomTypeValidator(NotNull))
	govalidator.CustomTypeTagMap.Set("phone", govalidator.CustomTypeValidator(Phone))
}

// ScanValues 将url.Values 数据解析到指定对象
func ScanValues(dst interface{}, v url.Values) error {
	if err := valuerDecoder.Decode(dst, v); nil != err {
		return errors.Wrapf(err, "解析url.Values数据到结构体")
	}

	ok, err := govalidator.ValidateStruct(dst)
	if nil != err {
		return errors.Wrap(err, "验证url.Values数据类型")
	}
	if !ok {
		return errors.New("验证url.Values数据")
	}

	return nil
}

// ScanJSON 将io.Reader数据解析到指定结构体
func ScanJSON(dst interface{}, b io.Reader) error {
	decoder := json.NewDecoder(b)

	if err := decoder.Decode(dst); nil != err {
		return errors.Wrap(err, "解析body数据到结构体")
	}

	ok, err := govalidator.ValidateStruct(dst)
	if nil != err {
		return errors.Wrap(err, "验证body数据类型")
	}
	if !ok {
		return errors.New("数据类型格式不正确")
	}

	return nil
}
