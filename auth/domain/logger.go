package domain

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}
