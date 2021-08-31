package loggers

import (
	"context"
	"fmt"
	"go-cleanarchitecture/domains"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

var _ domains.Logger = ZapLogger{}

func NewZapLogger(configFilePath string) (ZapLogger, error) {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.CallerKey = "file"
	encoderConfig.MessageKey = "msg"
	encoderConfig.LevelKey = "level"
	encoderConfig.TimeKey = "time"

	config := zap.NewDevelopmentConfig()
	config.Encoding = "json"
	config.EncoderConfig = encoderConfig

	logger, err := config.Build()
	if err != nil {
		return ZapLogger{}, err
	}

	return ZapLogger{logger}, nil
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
