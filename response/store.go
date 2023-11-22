package response

import sandboxrpc "github.com/redsift/go-sandbox-rpc/v3/sandboxrpc"

type StoreWrite struct {
	Store string                 `json:"store" msgpack:"store"`
	Data  []sandboxrpc.StoreItem `json:"data" msgpack:"data"`
}

func (StoreWrite) Key() string { return "write" }

func init() {
	RegisterItem(&StoreWrite{})
}
