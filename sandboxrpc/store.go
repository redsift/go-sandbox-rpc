package sandboxrpc

type StoreItem struct {
	Key string `json:"key" msgpack:"key"`
	StoreValue
}

type StoreValue struct {
	Value []byte `json:"value" msgpack:"value"`
	TTL   uint32 `json:"ttl" msgpack:"ttl"`
	Epoch int64  `json:"epoch" msgpack:"epoch"`
}
