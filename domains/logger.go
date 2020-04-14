package domains

// ロガー実装の抽象

type Logger interface {
	Debug(msg string)
	Error(msg string)
	Info(msg string)
	Warn(msg string)
}
