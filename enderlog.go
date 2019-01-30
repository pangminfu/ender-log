package enderlog

import "log"

type EnderLogger struct {
	loggerRepo []LoggerRepository
	debugmode  bool
}

func New() *EnderLogger {
	defaultRepo := NewDefaultLoggerRepository()
	s := &EnderLogger{
		loggerRepo: nil,
		debugmode:  false,
	}
	s.Add(defaultRepo)
	return s
}

func (s *EnderLogger) SetDebugMode(debugmode bool) {
	s.debugmode = debugmode
}

func (s *EnderLogger) Debug(msg string) {
	log.Println(msg)
}

func (s *EnderLogger) Add(repo LoggerRepository) {
	s.loggerRepo = append([]LoggerRepository{repo}, s.loggerRepo...)
}

func (s *EnderLogger) Info(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Info(msg)
	}
}

func (s *EnderLogger) Warn(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Warn(msg)
	}
}

func (s *EnderLogger) Error(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Error(msg)
	}
}

func (s *EnderLogger) Fatal(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debugmode && logger.Debugable() == false {
			continue
		}
		logger.Fatal(msg)
	}
}
