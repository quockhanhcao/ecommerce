package logger

import (
	"github.com/quockhanhcao/ecommerce/pkg/settings"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(loggerConfig settings.LoggerConfig) *LoggerZap {
	// debug -> info -> warn -> error -> fatal -> panic
	logLevel := loggerConfig.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   loggerConfig.FilePath,
		MaxSize:    loggerConfig.MaxSize,
		MaxBackups: loggerConfig.MaxBackups,
		MaxAge:     loggerConfig.MaxAge,
		Compress:   loggerConfig.Compress,
	}
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(&hook),
		),
		level,
	)
	return &LoggerZap{zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))}
}

// format log
func getEncoderLog() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> timestamp
	config.TimeKey = "timestamp"
	// from info INFO
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(config)
}
