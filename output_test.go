package loadtest

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	sampleURL, _ := url.Parse("https://example.com")
	c = dummyClock{}

	for _, test := range []struct {
		name     string
		id       string
		duration time.Duration
		request  *http.Request
		response *http.Response
		expect   string
	}{
		{"Well formed input", "abc", 100 * time.Millisecond, &http.Request{Method: "POST", URL: sampleURL}, &http.Response{StatusCode: 200, ContentLength: 100}, `{"sequenceID":"abc","url":"https://example.com","method":"POST","status":200,"size":100,"timestamp":"0001-01-01T00:00:00Z","duration":100000000,"error":null}`},
		{"Missing/ no response", "abc", 100 * time.Millisecond, &http.Request{Method: "POST", URL: sampleURL}, nil, `{"sequenceID":"abc","url":"https://example.com","method":"POST","status":0,"size":0,"timestamp":"0001-01-01T00:00:00Z","duration":100000000,"error":null}`},
	} {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if err != nil {
					t.Errorf("unexpected error %+v", err)
				}
			}()

			o := Parse(test.id, test.duration, test.request, test.response)
			if test.expect != o.String() {
				t.Errorf("expected `%s`, received `%s`", test.expect, o.String())
			}
		})
	}
}
