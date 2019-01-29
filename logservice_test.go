package logger

import "testing"

func TestLogI(t *testing.T) {
	slack := NewSlackLogger()
	svc := NewLogService()
	svc.add(slack)

	svc.i("test")
}
