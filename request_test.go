package golo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

type dummyClient struct {
	err bool
}

func (c *dummyClient) Do(req *http.Request) (r *http.Response, err error) {
	if c.err {
		err = fmt.Errorf("an error")
	}

	r = &http.Response{StatusCode: 200, Request: req, Body: ioutil.NopCloser(bytes.NewBufferString("hello world"))}

	return
}

func TestDoRequest(t *testing.T) {
	// This is a pretty crap test- we're just waiting to
	// see if it blows up: the most important part is what
	// is printed, but because we do this in a gofunc there's
	// no good way of capturing it as an example- we're always
	// going to be stuck either time.Sleep'ing, or hoping to
	// lose the race condition some other way.
	//
	// Instead: we can trust that so long as we get no panics,
	// and our tests around Parse() and Output.String() all
	// pass that this is working too.

	sampleURL, _ := url.Parse("https://example.com")

	for _, test := range []struct {
		name         string
		client       HTTPClient
		id           string
		req          *http.Request
		expectStatus int
	}{
		{"happy path", &dummyClient{}, "abc", &http.Request{URL: sampleURL}, 200},
		{"erroring client", &dummyClient{true}, "abc", &http.Request{URL: sampleURL}, 200},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("unexpected error %+v", err)
				}
			}()

			Client = test.client

			resp := DoRequest(test.id, test.req)
			if test.expectStatus != resp.StatusCode {
				t.Errorf("expected %d, received %d", test.expectStatus, resp.StatusCode)
			}
		})
	}
}

func TestNewSequenceID(t *testing.T) {
	id := NewSequenceID()
	if id == "" || id == DefaultSequenceID {
		t.Errorf("expected a random V4 uuid, received %q", id)
	}
}
