package logger

import "github.com/kataras/golog"

type DefaultLogger struct{}

func Default() *DefaultLogger {
	return &DefaultLogger{}
}

func (l *DefaultLogger) Info(v ...interface{}) {
	golog.Info(v...)
}

func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	golog.Infof(format, v...)
}

func (l *DefaultLogger) Error(v ...interface{}) {
	golog.Error(v...)
}

func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	golog.Errorf(format, v...)
}

func (l *DefaultLogger) Debug(v ...interface{}) {
	golog.Debug(v...)
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	golog.Debugf(format, v...)
}

func (l *DefaultLogger) Fatal(v ...interface{}) {
	golog.Fatal(v...)
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	golog.Fatalf(format, v...)
}
