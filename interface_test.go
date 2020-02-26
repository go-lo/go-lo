package golo

import (
	"context"
	"testing"
)

type dummyRunner struct {
	err bool
}

func (d dummyRunner) Trigger(c *Context, r *Response) (*Response, error) {
	r.Id = NewSequenceID()

	if d.err {
		r.Error = true
		r.Output = "An error"
	}

	tags := map[string]interface{}{
		"statuscode": 200,
	}

	r.Tags = Tagify(tags)

	return r, nil
}

func TestNew(t *testing.T) {
	r := dummyRunner{}

	_, err := New(r.Trigger)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}
}

func TestTrigger(t *testing.T) {
	r := dummyRunner{true}

	l, err := New(r.Trigger)
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	response, err := l.Trigger(context.Background(), &Context{JobName: "tests"})
	if err != nil {
		t.Errorf("unexpected error: %+v", err)
	}

	t.Run("Errors", func(t *testing.T) {
		if !response.Error {
			t.Errorf("expected Error to be set")
		}

		expect := "An error"
		if expect != response.Output {
			t.Errorf("expected %q, received %q", expect, response.Output)
		}
	})

	t.Run("Tags", func(t *testing.T) {
		if len(response.Tags) != 1 {
			t.Errorf("Expected 1 tags, recveived %d", len(response.Tags))
		}

		t.Run("Non-string values", func(t *testing.T) {
			expect := "200"
			received := response.Tags[0].Value

			if expect != received {
				t.Errorf("expected %q, received %q", expect, received)
			}
		})
	})
}
