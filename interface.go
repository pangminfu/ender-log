package enderlog

type LoggerRepository interface {
	SetDebugable(bool)
	Debugable() bool
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}

type LoggerService interface {
	Add(LoggerRepository)
	SetDebugMode(bool)
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
}
