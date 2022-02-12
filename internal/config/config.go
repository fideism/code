package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Load 导入
func Load(f string) {
	if _, err := toml.DecodeFile(f, &Setting); nil != err {
		panic(fmt.Sprintf("配置文件读取错误:%s", err.Error()))
	}
}
