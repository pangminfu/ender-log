package enderlog

import "testing"

func TestAdd_TwoLogRepo(t *testing.T) {
	log := New()
	log.Add(NewDefaultLoggerRepository())
	log.Add(NewDefaultSlackRepository("test"))

	if len(log.loggerRepo) != 2 {
		t.FailNow()
	}
}

func TestNew_NilRepo(t *testing.T) {
	log := New()

	if log.loggerRepo != nil {
		t.FailNow()
	}
}

func TestSetDebugMode_True(t *testing.T) {
	log := New()
	log.SetDebugMode(true)

	if log.debugmode != true {
		t.FailNow()
	}
}

func TestSetDebugMode_False(t *testing.T) {
	log := New()
	log.SetDebugMode(false)

	if log.debugmode != false {
		t.FailNow()
	}
}

func TestDefaultDebugMode_False(t *testing.T) {
	log := New()

	if log.debugmode != false {
		t.FailNow()
	}
}
