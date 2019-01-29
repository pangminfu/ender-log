package enderlog

import "log"

type DefaultLoggerRepository struct {
	debugable bool
}

func NewDefaultLoggerRepository() LoggerRepository {
	return &DefaultLoggerRepository{
		debugable: true,
	}
}

func (r *DefaultLoggerRepository) SetDebugable(debugable bool) {
	r.debugable = debugable
}

func (r *DefaultLoggerRepository) Debugable() bool {
	return r.debugable
}

func (r *DefaultLoggerRepository) Info(msg string) {
	log.Println("Info : " + msg)
}

func (r *DefaultLoggerRepository) Warn(msg string) {
	log.Println("Warn : " + msg)
}
func (r *DefaultLoggerRepository) Error(msg string) {
	log.Println("Error : " + msg)
}

func (r *DefaultLoggerRepository) Fatal(msg string) {
	log.Panicln("Fatal : " + msg)
}
