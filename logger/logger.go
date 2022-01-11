package logger

import (
	"go.uber.org/zap"
)

//=============================================
//日志实例，必须实现interface所有方法
//对外只暴露 Entity 实例， !!!! 项目中不在引入其他的日志工具
// 所有用到源日志pkg的原生方法，必须在此处重写
//=============================================
// 强制要求该接口方法必须实现
var _ Interface = (*Entity)(nil)

// Interface log interface
type Interface interface {
	Debug(msg string)
	Warn(msg string)
	Info(msg string)
	Error(msg string)
	Panic(msg string)
	WithError(err error) *Entity
	WithFields(key string, fields Fields) *Entity
}

// Config 日志配置信息
type Config struct {
	Channel    string
	Level      Level
	OutputFile string
	WithStack  bool
}

// Level 日志级别
type Level int8

const (
	// Debug debug
	Debug Level = iota + 1
	// Info Info
	Info
	// Warn Warn
	Warn
	// Error Error
	Error
	// Panic Panic
	Panic
)

// Entity 日志实例
type Entity struct {
	zap    *zap.Logger
	config Config
}

// NewLoggerEntity 实例化日志
func NewLoggerEntity(cfg Config) (*Entity, error) {
	z, err := newZap(cfg)
	if err != nil {
		return nil, err
	}

	return &Entity{
		zap:    z,
		config: cfg,
	}, nil
}
