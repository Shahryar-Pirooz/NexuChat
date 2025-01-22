package logger

import (
	"go.uber.org/zap"
)

func NewLogger(opt ...zap.Option) (*zap.Logger, error) {
	logger, err := zap.NewProduction(opt...)
	if err != nil {
		return nil, err
	}
	return logger, nil
}
