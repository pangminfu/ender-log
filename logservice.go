package logger

import "log"

type Service struct {
	loggerRepo []LoggerRepository
	debug      bool
}

func NewLogService() LoggerService {
	defaultRepo := NewDefaultLoggerRepository()
	s := &Service{
		loggerRepo: nil,
		debug:      false,
	}
	s.Add(defaultRepo)
	return s
}

func (s *Service) SetDebug(flag bool) {
	s.debug = flag
}

func (s *Service) Debug(msg string) {
	log.Println(msg)
}

func (s *Service) Add(repo LoggerRepository) {
	s.loggerRepo = append(s.loggerRepo, repo)
}

func (s *Service) Info(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debug && logger.Debugable() == false {
			continue
		}
		logger.Info(msg)
	}
}

func (s *Service) Warn(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debug && logger.Debugable() == false {
			continue
		}
		logger.Warn(msg)
	}
}

func (s *Service) Error(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debug && logger.Debugable() == false {
			continue
		}
		logger.Error(msg)
	}
}

func (s *Service) Fatal(msg string) {
	for _, logger := range s.loggerRepo {
		if s.debug && logger.Debugable() == false {
			continue
		}
		logger.Fatal(msg)
	}
}
