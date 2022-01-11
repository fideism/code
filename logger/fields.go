package logger

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Fields 日志使用
type Fields map[string]interface{}

// MakeFields 生成Fields
func MakeFields(data interface{}) Fields {
	p := make(Fields)

	j, e := json.Marshal(data)
	if e != nil {
		return p
	}

	if err := json.Unmarshal(j, &p); nil != err {
		return Fields{}
	}

	return p
}

// Set 设置值
func (f Fields) Set(k string, v interface{}) Fields {
	f[k] = v

	return f
}

// Get 获取值
func (f Fields) Get(k string) (v interface{}) {
	if v, ok := f[k]; ok {
		return v
	}

	return nil
}

// Exists 判断是否存在
func (f Fields) Exists(k string) bool {
	_, ok := f[k]

	return ok
}

// JSON 获取json
func (f Fields) JSON() string {
	j, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}

	return string(j)
}

// Value 返回数据库可识别类型
func (f Fields) Value() (driver.Value, error) {
	return json.Marshal(f)
}

// Scan ...
func (f *Fields) Scan(src interface{}) error {
	bytes, ok := src.([]byte)
	if !ok {
		return errors.New("util params value")
	}

	if err := json.Unmarshal(bytes, f); nil != err {
		return errors.New("fields value error")
	}

	return nil
}
