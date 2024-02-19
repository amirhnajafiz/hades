package logger

import (
	"github.com/go-logr/logr"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func New(level string) logr.Logger {
	opts := zap.Options{
		Development: true,
		Level:       labelToLevel(level),
	}

	return zap.New(zap.UseFlagOptions(&opts))
}

func labelToLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "error":
		return zapcore.ErrorLevel
	case "warn":
		return zapcore.WarnLevel
	case "info":
		return zapcore.InfoLevel
	default:
		return zapcore.PanicLevel
	}
}
