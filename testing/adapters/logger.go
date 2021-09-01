package adapters

import (
	"context"
	"go-cleanarchitecture/domains"
	"testing"
)

type MockLogger struct {
	T *testing.T
}

var _ domains.Logger = MockLogger{}

func (m MockLogger) Debugf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Debug] %s", msg)
	}
}

func (m MockLogger) Errorf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Error] %s", msg)
	}
}

func (m MockLogger) Infof(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Info] %s", msg)
	}
}

func (m MockLogger) Warnf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Warn] %s", msg)
	}
}
