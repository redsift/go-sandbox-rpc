package sandboxrpc

import "context"

type StreamSource interface {
	ID() string
	Ack(context.Context, string)
	Get(context.Context) (chan StreamItem, error)
}

type StreamSink interface {
	ID() string
	Put(context.Context, StreamItem) error
}

type StreamItem struct {
	Key     string
	Payload []byte
}

type StreamSpec struct {
	URL   string
	Token string
}
