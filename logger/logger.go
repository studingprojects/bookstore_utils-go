package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel   = "LOG_LEVEL"
	envOutputPath = "LOG_OUTPUT_PATH"
)

var (
	log logger
)

type bookstoreLogger interface {
	Printf(string, ...interface{})
	Print(...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{getOutputPath()},
		Level:       zap.NewAtomicLevelAt(getLevel()),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseColorLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger() bookstoreLogger {
	return log
}

func getLevel() zapcore.Level {
	switch os.Getenv(envLogLevel) {
	case "debug":
		return zap.DebugLevel
	case "error":
		return zap.ErrorLevel
	case "info":
		return zap.InfoLevel
	default:
		return zap.InfoLevel
	}
}

func getOutputPath() string {
	output := os.Getenv(envOutputPath)
	if output == "" {
		return "stdout"
	}
	return output
}

func (l logger) Printf(msg string, values ...interface{}) {
	if len(values) == 0 {
		Info(msg)
		return
	}
	Info(fmt.Sprintf(msg, values...))
}

func (l logger) Print(values ...interface{}) {
	Info(fmt.Sprintf("%v", values))
}

func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
