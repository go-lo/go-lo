[![Go Report Card](https://goreportcard.com/badge/github.com/go-lo/go-lo)](https://goreportcard.com/report/github.com/go-lo/go-lo)
[![Build Status](https://travis-ci.com/go-lo/go-lo.svg?branch=master)](https://travis-ci.com/go-lo/go-lo)
[![GoDoc](https://godoc.org/github.com/go-lo/go-lo?status.svg)](https://godoc.org/github.com/go-lo/go-lo)


# golo
`import "github.com/go-lo/go-lo"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package golo is a framework for running and writing distributed loadtests with go.

It does this by wrapping a special function, with the signature:


	func(*golo.Context, *golo.Response) (*golo.Response, error)

This function can then perform whatever tests it needs, returning loadtest run data which is then used
to report and chart on test runs.

A simple go-lo loadtest binary could be as simple as:


	package main
	
	import (
	    "net/http"
	
	    "github.com/go-lo/go-lo"
	)
	
	var (
	    url = "<a href="https://example.com">https://example.com</a>"
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

This loadtest can be uplaoded to a go-lo agent, with a schedule, and you should see results.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func NewSequenceID() string](#NewSequenceID)
* [func RegisterJobServer(s *grpc.Server, srv JobServer)](#RegisterJobServer)
* [type Context](#Context)
  * [func (*Context) Descriptor() ([]byte, []int)](#Context.Descriptor)
  * [func (m *Context) GetJobName() string](#Context.GetJobName)
  * [func (*Context) ProtoMessage()](#Context.ProtoMessage)
  * [func (m *Context) Reset()](#Context.Reset)
  * [func (m *Context) String() string](#Context.String)
  * [func (m *Context) XXX_DiscardUnknown()](#Context.XXX_DiscardUnknown)
  * [func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#Context.XXX_Marshal)
  * [func (m *Context) XXX_Merge(src proto.Message)](#Context.XXX_Merge)
  * [func (m *Context) XXX_Size() int](#Context.XXX_Size)
  * [func (m *Context) XXX_Unmarshal(b []byte) error](#Context.XXX_Unmarshal)
* [type JobClient](#JobClient)
  * [func NewJobClient(cc *grpc.ClientConn) JobClient](#NewJobClient)
* [type JobServer](#JobServer)
* [type Loadtest](#Loadtest)
  * [func New(f TriggerFunc) (l Loadtest, err error)](#New)
  * [func (l Loadtest) Start() (err error)](#Loadtest.Start)
  * [func (l Loadtest) Trigger(ctx context.Context, c *Context) (r *Response, err error)](#Loadtest.Trigger)
* [type Response](#Response)
  * [func (*Response) Descriptor() ([]byte, []int)](#Response.Descriptor)
  * [func (m *Response) GetError() bool](#Response.GetError)
  * [func (m *Response) GetId() string](#Response.GetId)
  * [func (m *Response) GetJobName() string](#Response.GetJobName)
  * [func (m *Response) GetOutput() string](#Response.GetOutput)
  * [func (m *Response) GetTags() []*ResponseTag](#Response.GetTags)
  * [func (*Response) ProtoMessage()](#Response.ProtoMessage)
  * [func (m *Response) Reset()](#Response.Reset)
  * [func (m *Response) String() string](#Response.String)
  * [func (m *Response) XXX_DiscardUnknown()](#Response.XXX_DiscardUnknown)
  * [func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#Response.XXX_Marshal)
  * [func (m *Response) XXX_Merge(src proto.Message)](#Response.XXX_Merge)
  * [func (m *Response) XXX_Size() int](#Response.XXX_Size)
  * [func (m *Response) XXX_Unmarshal(b []byte) error](#Response.XXX_Unmarshal)
* [type ResponseTag](#ResponseTag)
  * [func Tagify(m map[string]interface{}) (tags []*ResponseTag)](#Tagify)
  * [func (*ResponseTag) Descriptor() ([]byte, []int)](#ResponseTag.Descriptor)
  * [func (m *ResponseTag) GetKey() string](#ResponseTag.GetKey)
  * [func (m *ResponseTag) GetValue() string](#ResponseTag.GetValue)
  * [func (*ResponseTag) ProtoMessage()](#ResponseTag.ProtoMessage)
  * [func (m *ResponseTag) Reset()](#ResponseTag.Reset)
  * [func (m *ResponseTag) String() string](#ResponseTag.String)
  * [func (m *ResponseTag) XXX_DiscardUnknown()](#ResponseTag.XXX_DiscardUnknown)
  * [func (m *ResponseTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)](#ResponseTag.XXX_Marshal)
  * [func (m *ResponseTag) XXX_Merge(src proto.Message)](#ResponseTag.XXX_Merge)
  * [func (m *ResponseTag) XXX_Size() int](#ResponseTag.XXX_Size)
  * [func (m *ResponseTag) XXX_Unmarshal(b []byte) error](#ResponseTag.XXX_Unmarshal)
* [type TriggerFunc](#TriggerFunc)
* [type UnimplementedJobServer](#UnimplementedJobServer)
  * [func (*UnimplementedJobServer) Trigger(ctx context.Context, req *Context) (*Response, error)](#UnimplementedJobServer.Trigger)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/go-lo/go-lo/doc.go) [go-lo.pb.go](/src/github.com/go-lo/go-lo/go-lo.pb.go) [interface.go](/src/github.com/go-lo/go-lo/interface.go) [sequences.go](/src/github.com/go-lo/go-lo/sequences.go) [tags.go](/src/github.com/go-lo/go-lo/tags.go) 


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



## <a name="NewSequenceID">func</a> [NewSequenceID](/src/target/sequences.go?s=757:784#L21)
``` go
func NewSequenceID() string
```
NewSequenceID will return a fresh v4 uuid for sequences
of requests to use, to allow for ease of grouping journeys
together. This function swallows errors; should an error occur
then this will, instead, return loadtest.DefaultSequenceID.
Thus: a usable ID can always be guaranteed from this function



## <a name="RegisterJobServer">func</a> [RegisterJobServer](/src/target/go-lo.pb.go?s=8448:8501#L256)
``` go
func RegisterJobServer(s *grpc.Server, srv JobServer)
```



## <a name="Context">type</a> [Context](/src/target/go-lo.pb.go?s=745:995#L27)
``` go
type Context struct {
    JobName              string   `protobuf:"bytes,1,opt,name=jobName,proto3" json:"jobName,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

```









### <a name="Context.Descriptor">func</a> (\*Context) [Descriptor](/src/target/go-lo.pb.go?s=1159:1203#L37)
``` go
func (*Context) Descriptor() ([]byte, []int)
```



### <a name="Context.GetJobName">func</a> (\*Context) [GetJobName](/src/target/go-lo.pb.go?s=1809:1846#L59)
``` go
func (m *Context) GetJobName() string
```



### <a name="Context.ProtoMessage">func</a> (\*Context) [ProtoMessage](/src/target/go-lo.pb.go?s=1122:1152#L36)
``` go
func (*Context) ProtoMessage()
```



### <a name="Context.Reset">func</a> (\*Context) [Reset](/src/target/go-lo.pb.go?s=997:1022#L34)
``` go
func (m *Context) Reset()
```



### <a name="Context.String">func</a> (\*Context) [String](/src/target/go-lo.pb.go?s=1050:1083#L35)
``` go
func (m *Context) String() string
```



### <a name="Context.XXX_DiscardUnknown">func</a> (\*Context) [XXX_DiscardUnknown](/src/target/go-lo.pb.go?s=1667:1705#L53)
``` go
func (m *Context) XXX_DiscardUnknown()
```



### <a name="Context.XXX_Marshal">func</a> (\*Context) [XXX_Marshal](/src/target/go-lo.pb.go?s=1359:1434#L44)
``` go
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="Context.XXX_Merge">func</a> (\*Context) [XXX_Merge](/src/target/go-lo.pb.go?s=1500:1546#L47)
``` go
func (m *Context) XXX_Merge(src proto.Message)
```



### <a name="Context.XXX_Size">func</a> (\*Context) [XXX_Size](/src/target/go-lo.pb.go?s=1590:1622#L50)
``` go
func (m *Context) XXX_Size() int
```



### <a name="Context.XXX_Unmarshal">func</a> (\*Context) [XXX_Unmarshal](/src/target/go-lo.pb.go?s=1259:1306#L41)
``` go
func (m *Context) XXX_Unmarshal(b []byte) error
```



## <a name="JobClient">type</a> [JobClient](/src/target/go-lo.pb.go?s=7515:7630#L222)
``` go
type JobClient interface {
    Trigger(ctx context.Context, in *Context, opts ...grpc.CallOption) (*Response, error)
}
```
JobClient is the client API for Job service.

For semantics around ctx use and closing/ending streaming RPCs, please refer to <a href="https://godoc.org/google.golang.org/grpc#ClientConn.NewStream">https://godoc.org/google.golang.org/grpc#ClientConn.NewStream</a>.







### <a name="NewJobClient">func</a> [NewJobClient](/src/target/go-lo.pb.go?s=7680:7728#L230)
``` go
func NewJobClient(cc *grpc.ClientConn) JobClient
```




## <a name="JobServer">type</a> [JobServer](/src/target/go-lo.pb.go?s=8057:8140#L244)
``` go
type JobServer interface {
    Trigger(context.Context, *Context) (*Response, error)
}
```
JobServer is the server API for Job service.










## <a name="Loadtest">type</a> [Loadtest](/src/target/interface.go?s=490:557#L24)
``` go
type Loadtest struct {
    // contains filtered or unexported fields
}

```
Loadtest holds configuration and gRPC contexts which Loadtests
must be wrapped in







### <a name="New">func</a> [New](/src/target/interface.go?s=781:828#L33)
``` go
func New(f TriggerFunc) (l Loadtest, err error)
```
New takes scheduler code which implements the Runner
interface and returns a Server. It also runs some bootstrap
tasks to ensure a server has various things set that it
ought to, like a clock and an HTTPClient





### <a name="Loadtest.Start">func</a> (Loadtest) [Start](/src/target/interface.go?s=1046:1083#L46)
``` go
func (l Loadtest) Start() (err error)
```
Start will start an RPC server on loadtest.RPCAddr
and register Server ahead of Agents scheduling jobs




### <a name="Loadtest.Trigger">func</a> (Loadtest) [Trigger](/src/target/interface.go?s=1311:1394#L60)
``` go
func (l Loadtest) Trigger(ctx context.Context, c *Context) (r *Response, err error)
```
Trigger creates contexts/ outputs, and passes them to
(Loadtest).trigger() to run a test




## <a name="Response">type</a> [Response](/src/target/go-lo.pb.go?s=3299:3981#L113)
``` go
type Response struct {
    Id                   string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    JobName              string         `protobuf:"bytes,2,opt,name=jobName,proto3" json:"jobName,omitempty"`
    Error                bool           `protobuf:"varint,3,opt,name=error,proto3" json:"error,omitempty"`
    Output               string         `protobuf:"bytes,4,opt,name=output,proto3" json:"output,omitempty"`
    Tags                 []*ResponseTag `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
    XXX_NoUnkeyedLiteral struct{}       `json:"-"`
    XXX_unrecognized     []byte         `json:"-"`
    XXX_sizecache        int32          `json:"-"`
}

```









### <a name="Response.Descriptor">func</a> (\*Response) [Descriptor](/src/target/go-lo.pb.go?s=4149:4194#L127)
``` go
func (*Response) Descriptor() ([]byte, []int)
```



### <a name="Response.GetError">func</a> (\*Response) [GetError](/src/target/go-lo.pb.go?s=4985:5019#L163)
``` go
func (m *Response) GetError() bool
```



### <a name="Response.GetId">func</a> (\*Response) [GetId](/src/target/go-lo.pb.go?s=4811:4844#L149)
``` go
func (m *Response) GetId() string
```



### <a name="Response.GetJobName">func</a> (\*Response) [GetJobName](/src/target/go-lo.pb.go?s=4893:4931#L156)
``` go
func (m *Response) GetJobName() string
```



### <a name="Response.GetOutput">func</a> (\*Response) [GetOutput](/src/target/go-lo.pb.go?s=5074:5111#L170)
``` go
func (m *Response) GetOutput() string
```



### <a name="Response.GetTags">func</a> (\*Response) [GetTags](/src/target/go-lo.pb.go?s=5164:5207#L177)
``` go
func (m *Response) GetTags() []*ResponseTag
```



### <a name="Response.ProtoMessage">func</a> (\*Response) [ProtoMessage](/src/target/go-lo.pb.go?s=4111:4142#L126)
``` go
func (*Response) ProtoMessage()
```



### <a name="Response.Reset">func</a> (\*Response) [Reset](/src/target/go-lo.pb.go?s=3983:4009#L124)
``` go
func (m *Response) Reset()
```



### <a name="Response.String">func</a> (\*Response) [String](/src/target/go-lo.pb.go?s=4038:4072#L125)
``` go
func (m *Response) String() string
```



### <a name="Response.XXX_DiscardUnknown">func</a> (\*Response) [XXX_DiscardUnknown](/src/target/go-lo.pb.go?s=4666:4705#L143)
``` go
func (m *Response) XXX_DiscardUnknown()
```



### <a name="Response.XXX_Marshal">func</a> (\*Response) [XXX_Marshal](/src/target/go-lo.pb.go?s=4352:4428#L134)
``` go
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="Response.XXX_Merge">func</a> (\*Response) [XXX_Merge](/src/target/go-lo.pb.go?s=4495:4542#L137)
``` go
func (m *Response) XXX_Merge(src proto.Message)
```



### <a name="Response.XXX_Size">func</a> (\*Response) [XXX_Size](/src/target/go-lo.pb.go?s=4587:4620#L140)
``` go
func (m *Response) XXX_Size() int
```



### <a name="Response.XXX_Unmarshal">func</a> (\*Response) [XXX_Unmarshal](/src/target/go-lo.pb.go?s=4250:4298#L131)
``` go
func (m *Response) XXX_Unmarshal(b []byte) error
```



## <a name="ResponseTag">type</a> [ResponseTag](/src/target/go-lo.pb.go?s=1900:2243#L66)
``` go
type ResponseTag struct {
    Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
    Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
    XXX_NoUnkeyedLiteral struct{} `json:"-"`
    XXX_unrecognized     []byte   `json:"-"`
    XXX_sizecache        int32    `json:"-"`
}

```






### <a name="Tagify">func</a> [Tagify](/src/target/tags.go?s=138:197#L9)
``` go
func Tagify(m map[string]interface{}) (tags []*ResponseTag)
```
Tagify takes a map of strings to interfaces, and
turns it into tags which can be used in Responses





### <a name="ResponseTag.Descriptor">func</a> (\*ResponseTag) [Descriptor](/src/target/go-lo.pb.go?s=2423:2471#L77)
``` go
func (*ResponseTag) Descriptor() ([]byte, []int)
```



### <a name="ResponseTag.GetKey">func</a> (\*ResponseTag) [GetKey](/src/target/go-lo.pb.go?s=3121:3158#L99)
``` go
func (m *ResponseTag) GetKey() string
```



### <a name="ResponseTag.GetValue">func</a> (\*ResponseTag) [GetValue](/src/target/go-lo.pb.go?s=3208:3247#L106)
``` go
func (m *ResponseTag) GetValue() string
```



### <a name="ResponseTag.ProtoMessage">func</a> (\*ResponseTag) [ProtoMessage](/src/target/go-lo.pb.go?s=2382:2416#L76)
``` go
func (*ResponseTag) ProtoMessage()
```



### <a name="ResponseTag.Reset">func</a> (\*ResponseTag) [Reset](/src/target/go-lo.pb.go?s=2245:2274#L74)
``` go
func (m *ResponseTag) Reset()
```



### <a name="ResponseTag.String">func</a> (\*ResponseTag) [String](/src/target/go-lo.pb.go?s=2306:2343#L75)
``` go
func (m *ResponseTag) String() string
```



### <a name="ResponseTag.XXX_DiscardUnknown">func</a> (\*ResponseTag) [XXX_DiscardUnknown](/src/target/go-lo.pb.go?s=2967:3009#L93)
``` go
func (m *ResponseTag) XXX_DiscardUnknown()
```



### <a name="ResponseTag.XXX_Marshal">func</a> (\*ResponseTag) [XXX_Marshal](/src/target/go-lo.pb.go?s=2635:2714#L84)
``` go
func (m *ResponseTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error)
```



### <a name="ResponseTag.XXX_Merge">func</a> (\*ResponseTag) [XXX_Merge](/src/target/go-lo.pb.go?s=2784:2834#L87)
``` go
func (m *ResponseTag) XXX_Merge(src proto.Message)
```



### <a name="ResponseTag.XXX_Size">func</a> (\*ResponseTag) [XXX_Size](/src/target/go-lo.pb.go?s=2882:2918#L90)
``` go
func (m *ResponseTag) XXX_Size() int
```



### <a name="ResponseTag.XXX_Unmarshal">func</a> (\*ResponseTag) [XXX_Unmarshal](/src/target/go-lo.pb.go?s=2527:2578#L81)
``` go
func (m *ResponseTag) XXX_Unmarshal(b []byte) error
```



## <a name="TriggerFunc">type</a> [TriggerFunc](/src/target/interface.go?s=339:400#L20)
``` go
type TriggerFunc func(*Context, *Response) (*Response, error)
```
TriggerFunc is a function which a loadtest calls. All
loadtests have to do is implement this function in a go
application.










## <a name="UnimplementedJobServer">type</a> [UnimplementedJobServer](/src/target/go-lo.pb.go?s=8228:8266#L249)
``` go
type UnimplementedJobServer struct {
}

```
UnimplementedJobServer can be embedded to have forward compatible implementations.










### <a name="UnimplementedJobServer.Trigger">func</a> (\*UnimplementedJobServer) [Trigger](/src/target/go-lo.pb.go?s=8268:8360#L252)
``` go
func (*UnimplementedJobServer) Trigger(ctx context.Context, req *Context) (*Response, error)
```







- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
