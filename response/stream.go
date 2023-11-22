package response

type StreamAck []string

func (StreamAck) Key() string { return "ack" }

type StreamPub struct {
	ID    string `json:"id" msgpack:"id"`
	Value []byte `json:"value" msgpack:"value"`
}

func (StreamPub) Key() string { return "push" }

func init() {
	RegisterItem(&StreamAck{})
	RegisterItem(&StreamPub{})
}
