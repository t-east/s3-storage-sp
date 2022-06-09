package log

import (
	"sp/src/config"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func NewLogger() (*zap.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	var err error

	if config.IsDevelopment() {
		logger, err = zap.NewDevelopment()
	} else if config.IsTest() {
		logger = zap.NewNop()
	} else {
		logger, err = zap.NewProduction()
	}

	return logger, err
}
