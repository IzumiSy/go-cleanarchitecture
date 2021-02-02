package loggers

import (
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

func (logger ZapLogger) Debug(msg string) {
	logger.logger.Debug(msg)
}

func (logger ZapLogger) Error(msg string) {
	logger.logger.Error(msg)
}

func (logger ZapLogger) Info(msg string) {
	logger.logger.Info(msg)
}

func (logger ZapLogger) Warn(msg string) {
	logger.logger.Warn(msg)
}
