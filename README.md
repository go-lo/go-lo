

# loadtest
`import "github.com/jspc/loadtest"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
loadtest is a framework for running distributed loadtesting with go.

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
	        URL: "<a href="http://10.50.0.4:8765">http://10.50.0.4:8765</a>",
	    }
	
	    server := loadtest.NewServer(m)
	
	    panic(loadtest.StartListener(server))
	}

The important steps are:


	seq := loadtest.NewSequenceID()

A sequence ID is a string- using the same ID for all requests in a sequence of calls (completely analagous to a User Journey, say) allows us to identify slow routes better


	_ = loadtest.DoRequest(seq, req)

This executes *http.Request `req` with a sequence ID. This returns an *http.Response, and outputs pertinent json to STDOUT for the agent to pickup


	server := loadtest.NewServer(m)
	panic(loadtest.StartListener(server))

This will take our implementation of the interface loadtest.Runner and start up the RPC listener




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func DoRequest(id string, req *http.Request) (response *http.Response)](#DoRequest)
* [func NewSequenceID() string](#NewSequenceID)
* [func StartListener(server Server) error](#StartListener)
* [type NullArg](#NullArg)
* [type Output](#Output)
  * [func Parse(id string, duration time.Duration, r *http.Response) (o Output)](#Parse)
  * [func (o Output) String() string](#Output.String)
* [type Runner](#Runner)
* [type Server](#Server)
  * [func NewServer(r Runner) Server](#NewServer)
  * [func (s Server) Run(_ *NullArg, _ *NullArg) error](#Server.Run)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/jspc/loadtest/doc.go) [interface.go](/src/github.com/jspc/loadtest/interface.go) [output.go](/src/github.com/jspc/loadtest/output.go) [request.go](/src/github.com/jspc/loadtest/request.go) 


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

## <a name="pkg-variables">Variables</a>
``` go
var (
    // Client can be overridden for when extra control
    // is warranted, such as with authorization, or
    // overriding TLS configuration
    Client = &http.Client{}
)
```
``` go
var (
    RPCAddr = "127.0.0.1:9999"
)
```


## <a name="DoRequest">func</a> [DoRequest](/src/target/request.go?s=1048:1118#L34)
``` go
func DoRequest(id string, req *http.Request) (response *http.Response)
```
DoRequest will take an ID and an http.Request, turn it into
an Output, and print that to STDOUT with all of the pieces taken
care of. The purpose of this is to capture additional information,
such as duration and Sequence IDs.
Rather than pushing the responsibility of outputting this data to the
writer of a schedule, this function removes that boilerplate by
doing it it's self.



## <a name="NewSequenceID">func</a> [NewSequenceID](/src/target/request.go?s=1627:1654#L55)
``` go
func NewSequenceID() string
```
NewSequenceID will return a fresh v4 uuid for sequences
of requests to use, to allow for ease of grouping journeys
together. This function swallows errors; should an error occur
then this will, instead, return loadtest.DefaultSequenceID.
Thus: a usable ID can always be guaranteed from this function



## <a name="StartListener">func</a> [StartListener](/src/target/interface.go?s=927:966#L45)
``` go
func StartListener(server Server) error
```
StartListener will start an RPC server on loadtest.RPCAddr
and register Server ahead of Agents scheduling jobs




## <a name="NullArg">type</a> [NullArg](/src/target/interface.go?s=200:221#L15)
``` go
type NullArg struct{}
```
NullArg is a set of args that don't do anything
but that can be put into rpc calls to aid readability










## <a name="Output">type</a> [Output](/src/target/output.go?s=324:685#L15)
``` go
type Output struct {
    SequenceID string        `json:"sequenceID"`
    URL        string        `json:"url"`
    Method     string        `json:"method"`
    Status     int           `json:"status"`
    Size       int64         `json:"size"`
    Timestamp  time.Time     `json:"timestamp"`
    Duration   time.Duration `json:"duration"`
    Error      error         `json:"error"`
}
```
Output is a normalised, enriched struct containing
results for a response which can be printed and picked
up by a loadtest agent.

This has a number of convenience functions hanging
off the back of it to remove boilerplate in schedule code







### <a name="Parse">func</a> [Parse](/src/target/output.go?s=888:962#L30)
``` go
func Parse(id string, duration time.Duration, r *http.Response) (o Output)
```
Parse takes a sequence ID, duration, and an http.Response
and pulls out the necessary data an Output type wants
The sequence ID is useful to be able to group requests
in a journey together





### <a name="Output.String">func</a> (Output) [String](/src/target/output.go?s=1317:1348#L52)
``` go
func (o Output) String() string
```
String outputs a marshal'd json string for the attached
Output. It swallows errors.




## <a name="Runner">type</a> [Runner](/src/target/interface.go?s=372:404#L20)
``` go
type Runner interface {
    Run()
}
```
Runner is the interface to implement in scheduler
code; it provides a single function `Run()` which
takes no arguments, and returns nothing










## <a name="Server">type</a> [Server](/src/target/interface.go?s=488:525#L26)
``` go
type Server struct {
    // contains filtered or unexported fields
}
```
Server will expose scheduler code over RPC for agents
to run and work with.







### <a name="NewServer">func</a> [NewServer](/src/target/interface.go?s=623:654#L32)
``` go
func NewServer(r Runner) Server
```
NewServer takes scheduler code which implements the Runner
interface and returns a Server





### <a name="Server.Run">func</a> (Server) [Run](/src/target/interface.go?s=726:775#L37)
``` go
func (s Server) Run(_ *NullArg, _ *NullArg) error
```
Run is the RPC interface into scheduler code








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
