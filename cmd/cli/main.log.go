package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
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

func getWriterSync() zapcore.WriteSyncer {
	// write to file
	file, err := os.OpenFile("./log/log.txt", os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
