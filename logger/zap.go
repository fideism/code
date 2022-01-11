package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Info ...
func (e *Entity) Info(msg string) {
	if e.config.WithStack {
		e.zap.Info(msg, stacktraceField())
	} else {
		e.zap.Info(msg)
	}
}

// Error ...
func (e *Entity) Error(msg string) {
	if e.config.WithStack {
		e.zap.Error(msg, stacktraceField())
	} else {
		e.zap.Error(msg)
	}
}

// Debug ...
func (e *Entity) Debug(msg string) {
	if e.config.WithStack {
		e.zap.Debug(msg, stacktraceField())
	} else {
		e.zap.Debug(msg)
	}
}

// Warn ...
func (e *Entity) Warn(msg string) {
	if e.config.WithStack {
		e.zap.Warn(msg, stacktraceField())
	} else {
		e.zap.Warn(msg)
	}
}

// Panic ...
func (e *Entity) Panic(msg string) {
	if e.config.WithStack {
		e.zap.Panic(msg, stacktraceField())
	} else {
		e.zap.Panic(msg)
	}
}

// WithFields 日志字段
func (e *Entity) WithFields(key string, fields Fields) *Entity {
	j, _ := json.Marshal(fields)

	return &Entity{
		zap: e.zap.With(zapcore.Field{
			Key:       key,
			Type:      zapcore.ByteStringType,
			Interface: j,
		}),
		config: e.config,
	}
}

// WithError 日志中追加error信息
func (e *Entity) WithError(err error) *Entity {
	fields := zapcore.Field{
		Key:       "catch_error",
		Type:      zapcore.ErrorType,
		Interface: err,
	}

	return &Entity{
		zap:    e.zap.With(fields),
		config: e.config,
	}
}

func stacktraceField() zap.Field {
	return zap.ByteString("stacktrace", debug.Stack())
}

func zapLevel(l Level) zapcore.Level {
	switch l {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Panic:
		return zapcore.PanicLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZap(cfg Config) (*zap.Logger, error) {
	file := cfg.OutputFile
	if len(cfg.OutputFile) == 0 {
		file = "/tmp/runtime.log"
	}

	hook, err := rotate.New(
		file+"_%Y%m%d",
		rotate.WithLinkName(file),
		rotate.WithMaxAge(time.Hour*24*30),
		rotate.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		return nil, fmt.Errorf("实例化日志错误:%s", err.Error())
	}

	encoderConfig := zapcore.EncoderConfig{
		CallerKey:     "file",
		StacktraceKey: "stacktrace",
		EncodeName:    zapcore.FullNameEncoder,
		TimeKey:       "runtime",
		LevelKey:      "log_level",
		NameKey:       "logger",
		MessageKey:    "message",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	var writer zapcore.WriteSyncer
	if len(cfg.OutputFile) == 0 {
		writer = zapcore.AddSync(os.Stdout)
	} else {
		writer = zapcore.AddSync(hook)
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapLevel(cfg.Level),
	)

	//// 设置初始化字段
	field := zap.Fields(zap.String("channel", cfg.Channel))

	// 构造日志
	return zap.New(core, field), nil
}
