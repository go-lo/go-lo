// +build !output

package golo

import (
	"testing"
)

func TestRealClock_Now(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("unexpected error %+v", err)
		}
	}()

	_ = realClock{}.now()
}
