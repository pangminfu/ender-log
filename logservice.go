package logger

type Service struct {
	logRepo []LoggerRepository
	debug   bool
}

func NewLogService() LoggerService {
	return &Service{
		logRepo: nil,
		debug:   false,
	}
}

func (s *Service) SetDebug(flag bool) {

}

func (s *Service) Debug(msg string) {

}

func (s *Service) Add(repo LoggerRepository) {
	s.logRepo = append(s.logRepo, repo)
}

func (s *Service) Info(msg string) {
	for _, log := range s.logRepo {
		log.Info("")
	}
}

func (s *Service) Warn(msg string) {

}

func (s *Service) Error(msg string) {

}

func (s *Service) Fatal(msg string) {

}
