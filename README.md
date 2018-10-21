[![Go Report Card](https://goreportcard.com/badge/github.com/go-lo/go-lo)](https://goreportcard.com/report/github.com/go-lo/go-lo)
[![Build Status](https://travis-ci.com/go-lo/go-lo.svg?branch=master)](https://travis-ci.com/go-lo/go-lo)
[![GoDoc](https://godoc.org/github.com/go-lo/go-lo?status.svg)](https://godoc.org/github.com/go-lo/go-lo)


# golo
`import "github.com/go-lo/go-lo"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
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
	        URL: "<a href="http://localhost:8765">http://localhost:8765</a>",
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




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func DoRequest(id string, req *http.Request) (response *http.Response)](#DoRequest)
* [func NewSequenceID() string](#NewSequenceID)
* [func Start(server Server) (err error)](#Start)
* [type HTTPClient](#HTTPClient)
* [type NullArg](#NullArg)
* [type Output](#Output)
  * [func Parse(id string, duration time.Duration, r *http.Request, resp *http.Response) (o Output)](#Parse)
  * [func (o Output) String() string](#Output.String)
* [type Runner](#Runner)
* [type Server](#Server)
  * [func New(r Runner) Server](#New)
  * [func (s Server) Run(_ *NullArg, _ *NullArg) error](#Server.Run)


#### <a name="pkg-files">Package files</a>
[clock.go](/src/github.com/go-lo/go-lo/clock.go) [doc.go](/src/github.com/go-lo/go-lo/doc.go) [interface.go](/src/github.com/go-lo/go-lo/interface.go) [logging.go](/src/github.com/go-lo/go-lo/logging.go) [output.go](/src/github.com/go-lo/go-lo/output.go) [request.go](/src/github.com/go-lo/go-lo/request.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    // DefaultSequenceID is a uuid which will be returned should uuid.NewV4
    // fail. It can be safely compared with whatever is returned from
    // loadtest.SequenceID()- this uuid is a v5 uuid in the DNS namespace
    // whereas SequenceID() returns a v4 uuid.
    // see script/make_uuid.go in source repo for more information.
    DefaultSequenceID = "c276c8c7-6fec-5aa9-b6bd-4de12a49a9bb"
)
```
``` go
const (
    // RPCAddr is the default host on which a schedule listens
    // and an agent connects to
    RPCAddr = "127.0.0.1:9999"
)
```



## <a name="DoRequest">func</a> [DoRequest](/src/target/request.go?s=1681:1751#L51)
``` go
func DoRequest(id string, req *http.Request) (response *http.Response)
```
DoRequest will take an ID and an http.Request, turn it into
an Output, and print that to STDOUT with all of the pieces taken
care of. The purpose of this is to capture additional information,
such as duration and Sequence IDs.
Rather than pushing the responsibility of outputting this data to the
writer of a schedule, this function removes that boilerplate by
doing it its self.



## <a name="NewSequenceID">func</a> [NewSequenceID](/src/target/request.go?s=2567:2594#L90)
``` go
func NewSequenceID() string
```
NewSequenceID will return a fresh v4 uuid for sequences
of requests to use, to allow for ease of grouping journeys
together. This function swallows errors; should an error occur
then this will, instead, return loadtest.DefaultSequenceID.
Thus: a usable ID can always be guaranteed from this function



## <a name="Start">func</a> [Start](/src/target/interface.go?s=1643:1680#L76)
``` go
func Start(server Server) (err error)
```
Start will start an RPC server on loadtest.RPCAddr
and register Server ahead of Agents scheduling jobs




## <a name="HTTPClient">type</a> [HTTPClient](/src/target/request.go?s=259:388#L15)
``` go
type HTTPClient interface {
    // Do tracks https://golang.org/pkg/net/http/#Client.Do
    Do(*http.Request) (*http.Response, error)
}
```
HTTPClient is an interface which exposes a simple
way of doing http calls. It can be overwritten for
Oauth, or other auth, or even to stub calls out in
testing


``` go
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
```









## <a name="NullArg">type</a> [NullArg](/src/target/interface.go?s=299:320#L18)
``` go
type NullArg struct{}
```
NullArg is a set of args that don't do anything
but that can be put into rpc calls to aid readability










## <a name="Output">type</a> [Output](/src/target/output.go?s=338:699#L19)
``` go
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
```
Output is a normalised, enriched struct containing
results for a response which can be printed and picked
up by a loadtest agent.

This has a number of convenience functions hanging
off the back of it to remove boilerplate in schedule code







### <a name="Parse">func</a> [Parse](/src/target/output.go?s=902:996#L34)
``` go
func Parse(id string, duration time.Duration, r *http.Request, resp *http.Response) (o Output)
```
Parse takes a sequence ID, duration, and an http.Response
and pulls out the necessary data an Output type wants
The sequence ID is useful to be able to group requests
in a journey together





### <a name="Output.String">func</a> (Output) [String](/src/target/output.go?s=1317:1348#L53)
``` go
func (o Output) String() string
```
String outputs a marshal'd json string for the attached
Output. It swallows errors.




## <a name="Runner">type</a> [Runner](/src/target/interface.go?s=471:503#L23)
``` go
type Runner interface {
    Run()
}
```
Runner is the interface to implement in scheduler
code; it provides a single function `Run()` which
takes no arguments, and returns nothing










## <a name="Server">type</a> [Server](/src/target/interface.go?s=587:624#L29)
``` go
type Server struct {
    // contains filtered or unexported fields
}
```
Server will expose scheduler code over RPC for agents
to run and work with.







### <a name="New">func</a> [New](/src/target/interface.go?s=848:873#L37)
``` go
func New(r Runner) Server
```
New takes scheduler code which implements the Runner
interface and returns a Server. It also runs some bootstrap
tasks to ensure a server has various things set that it
ought to, like a clock and an HTTPClient





### <a name="Server.Run">func</a> (Server) [Run](/src/target/interface.go?s=1361:1410#L61)
``` go
func (s Server) Run(_ *NullArg, _ *NullArg) error
```
Run is the RPC interface into scheduler code








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
