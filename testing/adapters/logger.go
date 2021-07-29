package adapters

import (
	"go-cleanarchitecture/domains"
	"testing"
)

type MockLogger struct {
	T *testing.T
}

var _ domains.Logger = MockLogger{}

func (m MockLogger) Debug(msg string) {
	if m.T != nil {
		m.T.Logf("[Debug] %s", msg)
	}
}

func (m MockLogger) Error(msg string) {
	if m.T != nil {
		m.T.Logf("[Error] %s", msg)
	}
}

func (m MockLogger) Info(msg string) {
	if m.T != nil {
		m.T.Logf("[Info] %s", msg)
	}
}

func (m MockLogger) Warn(msg string) {
	if m.T != nil {
		m.T.Logf("[Warn] %s", msg)
	}
}
