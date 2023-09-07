package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"lqlzzz/go-card-notes/global"
	"os"
	"strings"
)

// InitZap //
// initialize zap logger
func InitZap() *zap.Logger {
	// 实例化core
	cores := getZapCore()
	// 新建实例
	logger := zap.New(zapcore.NewTee(cores...))
	return logger
}

// getZapCore //
// 获取zap配置
func getZapCore() []zapcore.Core {
	// 声明变量
	cores := make([]zapcore.Core, 0, 7)

	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志记录中时间的格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	for level := getLogLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		))
	}

	return cores
}

// getLogLevel //
// 获取配置中对应的zap日志级别
func getLogLevel() zapcore.Level {
	// 判断配置中的日志打印级别
	global.GCN_CONFIG.Zap.Level = strings.ToLower(global.GCN_CONFIG.Zap.Level)
	levelStr := global.GCN_CONFIG.Zap.Level
	switch levelStr {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
