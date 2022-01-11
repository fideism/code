# 日志包

封装 [`uber zap`](https://github.com/uber-go/zap) 日志包

接口形式封装，对外不暴露内部包任何方法调用

### 配置

```
logger.Config {
    Channel    string   // 分组，项目标识
    Level      Level    // 日志级别
    OutputFile string   // 输出文件路径，不填写默认输出到 终端
    WithStack  bool     // 是否在日志中记录调用堆栈信息  


}

// 日志级别

- logger.Debug
- logger.Warn
- logger.Info
- logger.Error
- logger.Panic
```

### 使用方法

- 实例化日志
```
cfg := logger.Config{
    Channel: "test",
    Level:   logger.InfoLevel,
}

entity := logger.NewLoggerEntity(cfg)
```

- 可使用方法

```
entity.Debug(msg string)
entity.Warn(msg string)
entity.Info(msg string)
entity.Error(msg string)
entity.Panic(msg string)
entity.Fatal(msg string)
entity.WithError(err error) *Entity
entity.WithFields(key string, fields Fields) *Entity
```

### 使用示例

```
cfg := logger.Config{
    Channel: "test",
    Level:   logger.InfoLevel,
}

entity := logger.NewLoggerEntity(cfg)

//entity.Info("测试")

//entity.Error("错误")

entity.WithError(errors.New("测试错误")).Info("只是测试错误")

entity.WithFields("params", logger.Fields{
    "name":  "tom",
    "phone": "13800138000",
}).Info("测试使用fields")
```