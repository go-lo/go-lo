package golo

import (
	"time"
)

type clock interface {
	now() time.Time
}

type realClock struct{}

func (_ realClock) now() time.Time {
	return time.Now()
}

type dummyClock struct{}

func (_ dummyClock) now() time.Time {
	return time.Time{}
}
