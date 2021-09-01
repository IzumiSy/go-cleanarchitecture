package adapters

import (
	"context"
	"fmt"
	"go-cleanarchitecture/domains"
	"testing"
)

type MockLogger struct {
	T *testing.T
}

var _ domains.Logger = MockLogger{}

func (m MockLogger) Debugf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Debug] %s", fmt.Sprintf(msg, a...))
	}
}

func (m MockLogger) Errorf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Error] %s", fmt.Sprintf(msg, a...))
	}
}

func (m MockLogger) Infof(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Info] %s", fmt.Sprintf(msg, a...))
	}
}

func (m MockLogger) Warnf(_ context.Context, msg string, a ...interface{}) {
	if m.T != nil {
		m.T.Logf("[Warn] %s", fmt.Sprintf(msg, a...))
	}
}
