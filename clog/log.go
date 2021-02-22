package clog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// Very verbose messages for debugging specific issues
	LevelDebug = "debug"
	// Default log level, informational
	LevelInfo = "info"
	// Warnings are messages about possible issues
	LevelWarn = "warn"
	// Errors are messages about things we know are problems
	LevelError = "error"
)

var Int64 = zap.Int64
var Int32 = zap.Int32
var Int = zap.Int
var String = zap.String
var Any = zap.Any
var Bool = zap.Bool
var Duration = zap.Duration

type Logger struct {
	zap          *zap.Logger
	consoleLevel zap.AtomicLevel
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func makeEncoder(json bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	encoderConfig.TimeKey = "timestamp"

	if json {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// func Info(msg string) {
// 	fmt.Fprintf(os.Stderr, "string\n")
// }

func NewLogger() *Logger {
	cores := []zapcore.Core{}
	logger := &Logger{}

	writer := zapcore.Lock(os.Stderr)
	core := zapcore.NewCore(makeEncoder(true), writer, getZapLevel(LevelDebug))
	cores = append(cores, core)
	combinedCore := zapcore.NewTee(cores...)
	logger.zap = zap.New(combinedCore, zap.AddCaller())
	return logger
}

func (l *Logger) Sugar() *SugarLogger {
	return &SugarLogger{}
}

func (l *Logger) Debug(message string, fields ...Field) {
	l.zap.Debug(message, fields...)
}

func (l *Logger) Info(message string, fields ...Field) {
	l.zap.Info(message, fields...)
}

func (l *Logger) Warn(message string, fields ...Field) {
	l.zap.Warn(message, fields...)
}

func (l *Logger) Error(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}

func (l *Logger) Critical(message string, fields ...Field) {
	l.zap.Error(message, fields...)
}
