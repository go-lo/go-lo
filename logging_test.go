// +build output

package golo

import (
	"net/http"
	"net/url"
	"time"
)

type loggingClient struct{}

func (c *loggingClient) Do(req *http.Request) (r *http.Response, err error) {
	return &http.Response{StatusCode: 200, Request: req}, nil
}

// This test should be run manually and infrequently- the order of the lines output
// are not guaranteed; they rely on the order requests succeed. The purpose is to ensure
// that the output is made of unique lines
func ExampleLogLoop() {
	c = dummyClock{}
	Client = &loggingClient{}

	url1, _ := url.Parse("https://example.com/1")
	req1 := &http.Request{URL: url1}

	url2, _ := url.Parse("https://example.com/2")
	req2 := &http.Request{URL: url2}

	url3, _ := url.Parse("https://example.com/3")
	req3 := &http.Request{URL: url3}

	go logLoop()

	go DoRequest("req1", req1)
	go DoRequest("req2", req2)
	go DoRequest("req3", req3)

	// gross
	time.Sleep(500 * time.Millisecond)

	// Output:
	// {"sequenceID":"req2","url":"https://example.com/2","method":"","status":200,"size":0,"timestamp":"0001-01-01T00:00:00Z","duration":0,"error":null}
	// {"sequenceID":"req3","url":"https://example.com/3","method":"","status":200,"size":0,"timestamp":"0001-01-01T00:00:00Z","duration":0,"error":null}
	// {"sequenceID":"req1","url":"https://example.com/1","method":"","status":200,"size":0,"timestamp":"0001-01-01T00:00:00Z","duration":0,"error":null}
}
