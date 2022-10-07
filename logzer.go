package logzer

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(debug bool) *zap.SugaredLogger {
	loggingLevel := zap.InfoLevel
	if debug {
		loggingLevel = zap.DebugLevel
	}

	cfgEncoder := zap.NewProductionEncoderConfig()
	cfgProd := zap.NewProductionConfig()

	cfgEncoder.TimeKey = "ts"
	cfgEncoder.MessageKey = "message"
	cfgEncoder.LevelKey = "level"
	cfgEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
	cfgProd.Level = zap.NewAtomicLevelAt(loggingLevel)
	cfgProd.EncoderConfig = cfgEncoder
	cfgProd.OutputPaths = []string{"stdout"}

	logger, _ := cfgProd.Build()
	if logger != nil {
		defer func() {
			_ = logger.Sync()
		}()
	}
	return logger.Sugar()
}
