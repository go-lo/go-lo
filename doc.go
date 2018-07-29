/*
Package golo is a framework for running distributed loadtesting with go.

It consists of multiple agents which receive jobs from a scheduler, and send results to a collector.


An agent will:

 * Receive a compiled loadtest (see below)
 * Run this locally
 * Send requests to this binary over rpc
 * Stream STDOUT _from_ the binary and send data to a collector

A simple loadtest looks like:

 package main

 import (
     "log"
     "net/http"

     "github.com/jspc/loadtest"
 )

 type MagnumAPI struct {
     URL string
 }

 func (m MagnumAPI) Run() {
     req, err := http.NewRequest("GET", m.URL, nil)
     if err != nil {
         panic(err)
     }

     seq := loadtest.NewSequenceID()

     _ = loadtest.DoRequest(seq, req)
 }

 func main() {
     m := MagnumAPI{
         URL: "http://10.50.0.4:8765",
     }

     server := loadtest.NewServer(m)

     panic(loadtest.StartListener(server))
 }

The important steps are:

     seq := loadtest.NewSequenceID()

A sequence ID is a string- using the same ID for all requests in a sequence of calls (completely analogous to a User Journey, say) allows us to identify slow routes better


     _ = loadtest.DoRequest(seq, req)

This executes *http.Request `req` with a sequence ID. This returns an *http.Response, and outputs pertinent json to STDOUT for the agent to pickup


     server := loadtest.NewServer(m)
     panic(loadtest.StartListener(server))

This will take our implementation of the interface loadtest.Runner and start up the RPC listener

*/
package golo
