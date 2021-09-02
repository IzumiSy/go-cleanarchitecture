package loggers

import (
	"context"
	"fmt"
	"go-cleanarchitecture/domains"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

var _ domains.Logger = ZapLogger{}

func NewZapLogger() (ZapLogger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return ZapLogger{}, err
	}

	return ZapLogger{logger.Sugar()}, nil
}

func (logger ZapLogger) Debugf(_ context.Context, msg string, a ...interface{}) {
	logger.logger.Debug(fmt.Sprintf(msg, a...))
}

func (logger ZapLogger) Errorf(_ context.Context, msg string, a ...interface{}) {
	logger.logger.Error(fmt.Sprintf(msg, a...))
}

func (logger ZapLogger) Infof(_ context.Context, msg string, a ...interface{}) {
	logger.logger.Info(fmt.Sprintf(msg, a...))
}

func (logger ZapLogger) Warnf(_ context.Context, msg string, a ...interface{}) {
	logger.logger.Warn(fmt.Sprintf(msg, a...))
}
