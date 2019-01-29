package logger

import "log"

type SlackLoggerRepository struct {
	slack string
}

func NewSlackLogger() LoggerRepository {
	return &SlackLoggerRepository{}
}

func (s *SlackLoggerRepository) Debug(msg string) {
	log.Println("slack")
}

func (s *SlackLoggerRepository) Info(msg string) {
	log.Println("slack")
}

func (s *SlackLoggerRepository) Warn(msg string) {

}

func (s *SlackLoggerRepository) Error(msg string) {

}

func (s *SlackLoggerRepository) Fatal(msg string) {

}
