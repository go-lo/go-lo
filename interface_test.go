package loadtest

import (
	"testing"
)

type dummyRunner struct{}

func (r dummyRunner) Run() {}

func TestNewServer(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error %+v", err)
		}
	}()

	NewServer(dummyRunner{})
}

func TestServer_Run(t *testing.T) {
	s := NewServer(dummyRunner{})

	err := s.Run(&NullArg{}, &NullArg{})
	if err != nil {
		t.Errorf("unexpectd error %+v", err)
	}
}
