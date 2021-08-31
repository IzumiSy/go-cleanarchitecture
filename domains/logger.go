package domains

import "context"

// ロガー実装の抽象
type Logger interface {
	Debugf(ctx context.Context, msg string, a ...interface{})
	Errorf(ctx context.Context, msg string, a ...interface{})
	Infof(ctx context.Context, msg string, a ...interface{})
	Warnf(ctx context.Context, msg string, a ...interface{})
}
