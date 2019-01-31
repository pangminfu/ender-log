package enderlog

import "log"

type EnderLogger struct {
	loggerRepo []LoggerRepository
	debugmode  bool
}

func New() *EnderLogger {
	s := &EnderLogger{
		loggerRepo: nil,
		debugmode:  false,
	}
	return s
}

func (el *EnderLogger) SetDebugMode(debugmode bool) {
	el.debugmode = debugmode
}

func (el EnderLogger) Debug(msg string) {
	log.Println(msg)
}

// Add will add your implementation of Logger so that enderlog can broadcast the log to your custom logger.
// Please be mindful if your logger implementation contain log.Panicln() that will terminate excution.
// you may need to add it at last so that your other logger implementation will work before execution being terminate
func (el *EnderLogger) Add(repo LoggerRepository) {
	el.loggerRepo = append(el.loggerRepo, repo)
}

func (el EnderLogger) Info(msg string) {
	for _, logger := range el.loggerRepo {
		if el.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Info(msg)
	}
}

func (el EnderLogger) Warn(msg string) {
	for _, logger := range el.loggerRepo {
		if el.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Warn(msg)
	}
}

func (el EnderLogger) Error(msg string) {
	for _, logger := range el.loggerRepo {
		if el.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Error(msg)
	}
}

func (el EnderLogger) Fatal(msg string) {
	for _, logger := range el.loggerRepo {
		if el.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Fatal(msg)
	}
}
