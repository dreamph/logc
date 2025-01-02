package zap

import (
	"context"
	"github.com/dreamph/logc"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) logc.Logger {
	return &zapLogger{
		logger: logger,
	}
}

func NewLogger(option *logc.Options) logc.Logger {
	var ioWriter = &lumberjack.Logger{
		Filename:   option.FilePath,
		MaxSize:    getDefault(option.MaxSize, 10), // Default: 10 MB
		MaxBackups: getDefault(option.MaxBackups, 0),
		MaxAge:     getDefault(option.MaxAge, 10), // Default: 10 days
		LocalTime:  true,
		Compress:   option.Compress,
	}

	writeFile := zapcore.AddSync(ioWriter)
	writeStdout := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	//encoderConfig.FunctionKey = "func"

	logLevel := parseLogLevel(option.Level, zap.InfoLevel)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writeFile, writeStdout),
		logLevel,
	)
	zapLog := &zapLogger{
		logger: zap.New(
			core,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}
	logc.SetDefaultLogger(zapLog)
	return zapLog

}

func getDefault(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func parseLogLevel(textLogLevel string, defaultLogLevel zapcore.Level) zapcore.Level {
	logLevel, err := zapcore.ParseLevel(textLogLevel)
	if err != nil {
		return defaultLogLevel
	}
	return logLevel
}

func GetLogger(l logc.Logger) *zap.Logger {
	return l.(*zapLogger).logger
}

func ToZapFields(fields map[string]interface{}) []zap.Field {
	var zapFields []zap.Field
	if len(fields) == 0 {
		return zapFields
	}

	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return zapFields
}

func (l *zapLogger) WithLogger(ctx context.Context) logc.Logger {
	fields := logc.GetFields(ctx)
	if len(fields) == 0 {
		return l
	}

	zapFields := ToZapFields(logc.GetFields(ctx))
	if len(zapFields) == 0 {
		return l
	}

	return New(l.logger.With(zapFields...))
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.logger.Sugar().Debug(args)
}

func (l *zapLogger) Info(args ...interface{}) {
	l.logger.Sugar().Info(args)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.logger.Sugar().Warn(args)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.logger.Sugar().Error(args)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *zapLogger) Release() {
	_ = l.logger.Sync()
}
