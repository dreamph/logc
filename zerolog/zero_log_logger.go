package zerolog

import (
	"context"
	"fmt"
	"github.com/dreamph/logc"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type zeroLogLogger struct {
	logger *zerolog.Logger
}

func New(logger *zerolog.Logger) logc.Logger {
	return &zeroLogLogger{
		logger: logger,
	}
}

func NewLogger(option *logc.Options) logc.Logger {
	// Configure the log file writer
	ioWriter := &lumberjack.Logger{
		Filename:   option.FilePath,
		MaxSize:    getDefault(option.MaxSize, 10), // Default: 10 MB
		MaxBackups: getDefault(option.MaxBackups, 0),
		MaxAge:     getDefault(option.MaxAge, 10), // Default: 10 days
		LocalTime:  true,
		Compress:   option.Compress,
	}

	// Multi-writer for file and stdout
	multiWriter := io.MultiWriter(ioWriter, os.Stdout)

	// Initialize zerolog
	zerolog.TimeFieldFormat = time.RFC3339

	zeroLog := zerolog.New(multiWriter).
		Level(parseLevel(option.Level, zerolog.InfoLevel)).
		With().
		Timestamp().
		Caller().
		Stack().
		Logger()

	customLogger := &zeroLogLogger{
		logger: &zeroLog,
	}
	logc.SetDefaultLogger(customLogger)
	return customLogger

}

func getDefault(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func parseLevel(level string, defaultLevel zerolog.Level) zerolog.Level {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return defaultLevel
	}
	return logLevel
}

func (l *zeroLogLogger) Debug(args ...interface{}) {
	l.logger.Debug().Msg(getMessage(args))
}

func (l *zeroLogLogger) Info(args ...interface{}) {
	l.logger.Info().Msg(getMessage(args))
}

func (l *zeroLogLogger) Warn(args ...interface{}) {
	l.logger.Warn().Msg(getMessage(args))
}

func (l *zeroLogLogger) Error(args ...interface{}) {
	l.logger.Error().Msg(getMessage(args))
}

func (l *zeroLogLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msg(formatMessage(template, args))
}

func (l *zeroLogLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msg(formatMessage(template, args))
}

func (l *zeroLogLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msg(formatMessage(template, args))
}

func (l *zeroLogLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msg(formatMessage(template, args))
}

func GetLogger(l logc.Logger) *zerolog.Logger {
	return l.(*zeroLogLogger).logger
}

func (l *zeroLogLogger) Release() {
}

func (l *zeroLogLogger) WithLogger(ctx context.Context) logc.Logger {
	fields := logc.GetFields(ctx)
	if len(fields) == 0 {
		return l
	}

	zeroLog := l.logger.With().Fields(fields).Logger()
	return New(&zeroLog)
}

func getMessage(args []interface{}) string {
	if len(args) == 1 {
		if str, ok := args[0].(string); ok {
			return str
		}
	}
	return fmt.Sprint(args...)
}

func formatMessage(template string, args []interface{}) string {
	if len(args) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, args...)
	}
	return getMessage(args)
}

/*
type CallerHook struct {
	skip int
}

// Run adjusts the caller field in the log event
func (h CallerHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level != zerolog.NoLevel {
		_, file, line, ok := runtime.Caller(h.skip)
		if ok {
			e.Caller().Str("caller", fmt.Sprintf("%s:%d", file, line))
		}
	}
}*/
