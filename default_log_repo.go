package logger

import "log"

type DefaultLoggerRepository struct {
	debug bool
}

func NewDefaultLoggerRepository() LoggerRepository {
	return &DefaultLoggerRepository{
		debug: true,
	}
}

func (r *DefaultLoggerRepository) SetDebugable(debug bool) {
	r.debug = debug
}

func (r *DefaultLoggerRepository) Debugable() bool {
	return r.debug
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
