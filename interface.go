package logger

type LoggerRepository interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type LoggerService interface {
	Add(LoggerRepository)
	SetDebug(bool)
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}
