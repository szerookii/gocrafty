package logger

type Logger interface {
	SetLevel(level string)

	Info(v ...any)
	Error(v ...any)
	Debug(v ...any)
	Fatal(v ...any)

	Infof(format string, v ...any)
	Errorf(format string, v ...any)
	Debugf(format string, v ...any)
	Fatalf(format string, v ...any)
}
