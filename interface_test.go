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

func TestSetupListener(t *testing.T) {
	s := NewServer(dummyRunner{})

	t.Run("clean configuration", func(t *testing.T) {
		_, l, err := setupListener(s)
		defer l.Close()

		if err != nil {
			t.Errorf("unexpected error %+v", err)
		}
	})

	t.Run("trying to bind to used port", func(t *testing.T) {
		_, l, _ := setupListener(s)
		defer l.Close()

		_, _, err := setupListener(s)
		if err == nil {
			t.Errorf("expected error")
		}
	})
}
