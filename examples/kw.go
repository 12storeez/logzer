package main

import (
	"errors"
	"github.com/12storeez/logzer"
	"go.uber.org/zap"
)

func main() {
	// init logger with debug mode
	logger := logzer.New(true)

	logger.Debugw("logger.Debug()", zap.String("debug", "enabled"))
	logger.Debugw("logger.Debug()", zap.Bool("debug", true))
	logger.Infow("logger.Info()",
		zap.Int("client_id", 1),
		zap.String("client_phone", "+79010000001"))
	logger.Infow("logger.Info()", zap.Float64("price", 25.64))
	logger.Errorw("logger.Error()", zap.Error(errors.New("generated error")))

	// rewrite logger with one without debug mode
	logger = logzer.New(false)
	logger.Debugw("logger.Debug()", zap.String("debug", "disabled"))
	logger.Infow("logger.Info()", zap.Strings("emails", []string{"test@test.org", "example@test.org"}))
	logger.Warnw("logger.Warn()", zap.Ints("ids", []int{14562, 64567}))
}
