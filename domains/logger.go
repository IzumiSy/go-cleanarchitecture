package domains

type Logger interface {
	Debug(msg string)
	Error(msg string)
	Info(msg string)
	Warn(msg string)

	// Debugf(template string, args ...interface{})
	// Errorf(template string, args ...interface{})
	// Infof(template string, args ...interface{})
	// Warnf(template string, args ...interface{})
}
