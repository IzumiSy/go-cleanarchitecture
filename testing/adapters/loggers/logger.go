package loggers

import (
    "go-cleanarchitecture/domains"
)

type MockLogger struct{}

var _ domains.Logger = MockLogger{}

func (_ MockLogger) Debug(msg string) {}

func (_ MockLogger) Error(msg string) {}

func (_ MockLogger) Info(msg string) {}

func (_ MockLogger) Warn(msg string) {}
