package logger

type DefaultSlackRepository struct {
	slack string
}

func NewDefaultSlackRepository() LoggerRepository {
	return &DefaultSlackRepository{}
}

func (s *DefaultSlackRepository) Debug(msg string) {

}

func (s *DefaultSlackRepository) Info(msg string) {

}

func (s *DefaultSlackRepository) Warn(msg string) {

}

func (s *DefaultSlackRepository) Error(msg string) {

}

func (s *DefaultSlackRepository) Fatal(msg string) {

}
