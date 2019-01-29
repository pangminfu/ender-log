package logger

import "log"

type Service struct {
	loggerRepo []LoggerRepository
	debug      bool
}

func NewLogService() LoggerService {
	return &Service{
		loggerRepo: nil,
		debug:      false,
	}
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
	log.Println("Info : " + msg)
	if s.debug == true {
		return
	}

	for _, logger := range s.loggerRepo {
		logger.Info(msg)
	}
}

func (s *Service) Warn(msg string) {
	log.Println("Warn : " + msg)
	if s.debug == true {
		return
	}

	for _, logger := range s.loggerRepo {
		logger.Warn(msg)
	}
}

func (s *Service) Error(msg string) {
	log.Println("Error : " + msg)
	if s.debug == true {
		return
	}

	for _, logger := range s.loggerRepo {
		logger.Error(msg)
	}
}

func (s *Service) Fatal(msg string) {
	for _, logger := range s.loggerRepo {
		logger.Fatal(msg)
	}

	log.Panicln("Fatal : " + msg)
}
