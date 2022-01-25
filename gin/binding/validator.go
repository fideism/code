package binding

import (
	"regexp"
	"time"
)

var (
	rxPhone = regexp.MustCompile(`^\+?\d+$`)
)

// SliceString 字符串切片
// tagName: slice.string
func SliceString(i, o interface{}) bool {
	switch i.(type) {
	case []string:
		return true
	default:
		return false
	}
}

// NotNull 字符串不可为空
func NotNull(i, o interface{}) bool {
	switch v := i.(type) {
	case string:
		return len(v) != 0
	case []byte:
		return len(v) != 0
	case int64:
		return v != 0
	case int32:
		return v != 0
	case int16:
		return v != 0
	case int8:
		return v != 0
	case int:
		return v != 0
	case float64:
		return v != 0
	case float32:
		return v != 0
	case time.Time:
		var zeroTime time.Time
		return v.UnixNano() != zeroTime.UnixNano()
	}

	return false
}

// Phone 手机号判断
// 使用正则表达式判断，这里简单的判断 +861111111111
func Phone(i, o interface{}) bool {
	switch v := i.(type) {
	case string:
		return rxPhone.MatchString(v)
	case []byte:
		return rxPhone.MatchString(string(v))
	default:
		return false
	}
}
