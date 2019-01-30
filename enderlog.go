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

func (s *EnderLogger) SetDebugMode(debugmode bool) {
	s.debugmode = debugmode
}

func (s *EnderLogger) Debug(msg string) {
	log.Println(msg)
}

// Add will add your implementation of Logger so that enderlog can broadcast the log to your custom logger.
// Please be mindful if your logger implementation contain log.Panicln() that will terminate excution.
// you may need to add it at last so that your other logger implementation will work before execution being terminate
func (s *EnderLogger) Add(repo LoggerRepository) {
	s.loggerRepo = append(s.loggerRepo, repo)
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
