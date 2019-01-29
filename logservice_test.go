package logger

import "testing"

func TestLogI(t *testing.T) {
	slack := NewDefaultSlackRepository()
	svc := NewLogService()
	svc.Add(slack)

	svc.Info("test")
}
