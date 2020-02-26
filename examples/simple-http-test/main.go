package main

import (
	"net/http"

	"github.com/go-lo/go-lo"
)

var (
	url = "https://example.com"
)

func Trigger(c *golo.Context, r *golo.Response) (*golo.Response, error) {
	resp, err := http.Get(url)

	if err != nil {
		r.Error = true
		r.Output = err.Error()
	}

	// Set the Job ID for this run
	r.Id = golo.NewSequenceID()

	// Add some tags
	r.Tags = golo.Tagify(map[string]interface{}{
		"status": resp.Status,
		"size":   resp.ContentLength,
		"url":    url,
	})

	return r, nil
}

func main() {
	loadtest, err := golo.New(Trigger)
	if err != nil {
		panic(err)
	}

	err = loadtest.Start()
	if err != nil {
		panic(err)
	}
}
