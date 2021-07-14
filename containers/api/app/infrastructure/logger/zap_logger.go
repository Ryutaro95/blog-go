package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func init() {
	level := zap.NewAtomicLevel()
	var encoding string
	var encodeLevel zapcore.LevelEncoder

	switch os.Getenv("BACKEND_ENV") {
	case "production", "staging":
		level.SetLevel(zapcore.InfoLevel)
		encoding = "json"
		encodeLevel = zapcore.CapitalLevelEncoder
	default:
		level.SetLevel(zapcore.DebugLevel)
		encoding = "console"
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}

	myConfig := zap.Config{
		Level:    level,
		Encoding: encoding,
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "Time",
			LevelKey:       "Level",
			NameKey:        "Name",
			CallerKey:      "Caller",
			MessageKey:     "Msg",
			StacktraceKey:  "St",
			EncodeLevel:    encodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, _ := myConfig.Build(zap.AddCallerSkip(1))

	zapLogger = logger
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zapLogger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zapcore.Field) {
	zapLogger.Fatal(msg, fields...)
}
