package logger

type LoggerInterface interface {
	SetLevel(level int)
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Close()
}
