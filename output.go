package golo

import (
	"encoding/json"
	"net/http"
	"reflect"
	"time"
)

var (
	c clock

	errorType = reflect.TypeOf((*error)(nil)).Elem()
)

// Output is a normalised, enriched struct containing
// results for a response which can be printed and picked
// up by a loadtest agent.
//
// This has a number of convenience functions hanging
// off the back of it to remove boilerplate in schedule code
type Output struct {
	SequenceID string        `json:"sequenceID"`
	URL        string        `json:"url"`
	Method     string        `json:"method"`
	Status     int           `json:"status"`
	Size       int64         `json:"size"`
	Timestamp  time.Time     `json:"timestamp"`
	Duration   time.Duration `json:"duration"`
	Error      interface{}   `json:"error"`
}

// Parse takes a sequence ID, duration, and an http.Response
// and pulls out the necessary data an Output type wants
// The sequence ID is useful to be able to group requests
// in a journey together
func Parse(id string, duration time.Duration, r *http.Request, resp *http.Response) (o Output) {
	o = Output{
		SequenceID: id,
		Timestamp:  c.now(),
		Duration:   duration,
		URL:        r.URL.String(),
		Method:     r.Method,
	}

	if resp != nil {
		o.Status = resp.StatusCode
		o.Size = resp.ContentLength
	}

	return
}

// String outputs a marshal'd json string for the attached
// Output. It swallows errors.
func (o Output) String() string {
	if o.Error != nil {
		if reflect.TypeOf(o.Error).Implements(errorType) {
			o.Error = o.Error.(error).Error()
		}
	}

	output, _ := json.Marshal(o)

	return string(output)
}
