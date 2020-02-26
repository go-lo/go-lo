package golo

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
)

const (
	// RPCAddr is the default host on which a schedule listens
	// and an agent connects to
	RPCAddr = "127.0.0.1:9999"
)

// TriggerFunc is a function which a loadtest calls. All
// loadtests have to do is implement this function in a go
// application.
type TriggerFunc func(*Context, *Response) (*Response, error)

// Loadtest holds configuration and gRPC contexts which Loadtests
// must be wrapped in
type Loadtest struct {
	trigger TriggerFunc
	server  *grpc.Server
}

// New takes scheduler code which implements the Runner
// interface and returns a Server. It also runs some bootstrap
// tasks to ensure a server has various things set that it
// ought to, like a clock and an HTTPClient
func New(f TriggerFunc) (l Loadtest, err error) {
	l = Loadtest{
		trigger: f,
	}

	l.server = grpc.NewServer()
	RegisterJobServer(l.server, l)

	return
}

// Start will start an RPC server on loadtest.RPCAddr
// and register Server ahead of Agents scheduling jobs
func (l Loadtest) Start() (err error) {
	// Listen to new requests
	lis, err := net.Listen("tcp", RPCAddr)
	if err != nil {
		return
	}

	l.server.Serve(lis)

	return
}

// Trigger creates contexts/ outputs, and passes them to
// (Loadtest).trigger() to run a test
func (l Loadtest) Trigger(ctx context.Context, c *Context) (r *Response, err error) {
	r = &Response{
		JobName: c.JobName,
		Tags:    make([]*ResponseTag, 0),
	}

	start := time.Now()
	r, err = l.trigger(c, r)
	r.Duration = time.Now().Sub(start)

	return
}
