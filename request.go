package loadtest

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

// HTTPClient is an interface which exposes a simple
// way of doing http calls. It can be overwritten for
// Oauth, or other auth, or even to stub calls out in
// testing
type HTTPClient interface {
	// Do tracks https://golang.org/pkg/net/http/#Client.Do
	Do(*http.Request) (*http.Response, error)
}

var (
	// Client can be overridden for when extra control
	// is warranted, such as with authorization, or
	// overriding TLS configuration
	Client HTTPClient

	// CloseRequests will ensure all requests are closed
	// as early as possible, as if Keep Alive is disabled.
	// This defaults to true to:
	//  1. Ensure connections don't hang around slupring resources, and
	//  2. Because keep alive isn't necessarily a great way to prove the
	//     performance of an endpoint
	CloseRequests = true
)

const (
	// DefaultSequenceID is a uuid which will be returned should uuid.NewV4
	// fail. It can be safely compared with whatever is returned from
	// loadtest.SequenceID()- this uuid is a v5 uuid in the DNS namespace
	// whereas SequenceID() returns a v4 uuid.
	// see script/make_uuid.go in source repo for more information.
	DefaultSequenceID = "c276c8c7-6fec-5aa9-b6bd-4de12a49a9bb"
)

// DoRequest will take an ID and an http.Request, turn it into
// an Output, and print that to STDOUT with all of the pieces taken
// care of. The purpose of this is to capture additional information,
// such as duration and Sequence IDs.
// Rather than pushing the responsibility of outputting this data to the
// writer of a schedule, this function removes that boilerplate by
// doing it it's self.
func DoRequest(id string, req *http.Request) (response *http.Response) {
	if Client == nil {
		Client = &http.Client{}
	}

	if CloseRequests {
		req.Close = true
	}

	start := c.now()
	response, err := Client.Do(req)
	end := c.now()

	if err == nil {
		defer response.Body.Close()
	}

	o := Parse(id, end.Sub(start), req, response)

	if err != nil {
		o.Error = err
	}

	// Let this happen out of band- we've already done the
	// difficult stuff
	go fmt.Println(o.String())

	return
}

// NewSequenceID will return a fresh v4 uuid for sequences
// of requests to use, to allow for ease of grouping journeys
// together. This function swallows errors; should an error occur
// then this will, instead, return loadtest.DefaultSequenceID.
// Thus: a usable ID can always be guaranteed from this function
func NewSequenceID() string {
	s, err := uuid.NewV4()
	if err != nil {
		return DefaultSequenceID
	}

	return s.String()
}
