/*
Package golo is a framework for running distributed loadtesting with go.

When used in a loadtest it will:

 * Provide a make http requests using anything which implements the golo.HTTPClient interface (the default value is an *http.Client from "net/http")
 * Turn responses into the lineformat an agent (github.com/go-lo/agent) can parse and handle
 * Prints these lines to STDOUT (via a gofunc) where an agent can parse them (this is to put the burden of buffering, parsing, and collecting onto an agent to allow a loadtest schedule to concentrate on just making requests

It does all of this by exposing an RPC service which an agent uses to schedule a call.

This library, then, is most useful when used with the rest of the go-lo suite but can realisitically be used by anything which works like a go-lo agent.

A simple loadtest looks like:

 package main

 import (
     "log"
     "net/http"

     "github.com/go-lo/go-lo"
 )

 type API struct {
     URL string
 }

 func (a API) Run() {
     req, err := http.NewRequest("GET", m.URL, nil)
     if err != nil {
         panic(err)
     }

     seq := golo.NewSequenceID()

     _ = golo.DoRequest(seq, req)
 }

 func main() {
     a := API{
         URL: "http://localhost:8765",
     }

     server := golo.New(m)

     panic(golo.Start(server))
 }

The important steps are:

     seq := golo.NewSequenceID()

A sequence ID is a string- using the same ID for all requests in a sequence of calls (completely analogous to a User Journey, say) allows us to identify slow routes better


     _ = golo.DoRequest(seq, req)

This executes *http.Request `req` with a sequence ID. This returns an *http.Response, and outputs pertinent json to STDOUT for the agent to pickup


     server := golo.New(m)
     panic(golo.Start(server))

This will take our implementation of the interface golo.Runner and start up the RPC listener

*/
package golo
