package logc

import (
	"context"
	"log"
	"os"
)

var defaultLogger Logger

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})

	Release()

	WithLogger(ctx context.Context) Logger
}

type Options struct {
	FilePath   string
	Level      string
	Format     string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func GetLogger() Logger {
	return defaultLogger
}

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

func LogErrorAndExit(err error, appLogger ...Logger) {
	if len(appLogger) == 1 {
		appLogger[0].Error(err)
	} else {
		log.Println(err)
	}

	os.Exit(1)
}

type Key string

const loggerFieldKey = Key("loggerFieldKey")

func WithValue(ctx context.Context, fields map[string]interface{}) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	return context.WithValue(ctx, loggerFieldKey, fields)
}

func GetFields(ctx context.Context) map[string]interface{} {
	fields, ok := ctx.Value(loggerFieldKey).(map[string]interface{})
	if ok {
		return fields
	}
	return nil
}
