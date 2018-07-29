package golo

import (
	"time"
)

type clock interface {
	now() time.Time
}

type realClock struct{}

func (realClock) now() time.Time {
	return time.Now()
}

type dummyClock struct{}

func (dummyClock) now() time.Time {
	return time.Time{}
}
