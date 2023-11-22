package sandboxrpc

import "github.com/redsift/go-sandbox-rpc/v3/sandboxrpc"

type StoreItem = sandboxrpc.StoreItem
type StoreValue = sandboxrpc.StoreValue

type StoreData interface {
	Get() (chan StoreItem, error)
	Size() (size int, finite bool)
}

type InlineStoreData struct {
	Data map[string]StoreValue `json:"data" msgpack:"data"`
}

func (isd *InlineStoreData) Get() (chan StoreItem, error) {
	ch := make(chan StoreItem)
	go func() {
		for k, i := range isd.Data {
			ch <- StoreItem{Key: k, StoreValue: i}
		}
	}()
	return ch, nil
}

func (isd *InlineStoreData) Size() (int, bool) {
	return len(isd.Data), true
}

type StreamingStoreData struct {
	Stream chan StoreItem
}

func (ssd *StreamingStoreData) Get() (chan StoreItem, error) {
	return ssd.Stream, nil
}

func (ssd *StreamingStoreData) Size() (int, bool) {
	return len(ssd.Stream), false
}
