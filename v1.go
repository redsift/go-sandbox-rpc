// sandbox_rpc.go
// Sandbox RPC model
//
// Copyright (c) 2017 RedSift Limited. All rights reserved.
package sandboxrpc

// ComputeRequestV1 contains the parameters to invoke the node implementation function.
type ComputeRequestV1 struct {
	In    *StoredDataQuantum   `json:"in"`
	Query []string             `json:"query,omitempty"`
	With  *StoredDataQuantum   `json:"with,omitempty"`
	Get   []*StoredDataQuantum `json:"get,omitempty"`
}

// ResponseV1 is returned by the sandbox.
type ResponseV1 struct {
	Out   []*ComputeResponseV1 `json:"out,omitempty"`
	Error interface{}          `json:"error,omitempty"`
	Stats interface{}          `json:"stats,omitempty"` // timings and resource usage to be booked against this sift TODO: nail down
}

// A single unit of operation
type StoredDataQuantum struct {
	Bucket string        `json:"bucket"`
	Key    string        `json:"key,omitempty"`
	Data   []*StoredData `json:"data"`
}

type Data struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
	Epoch int64  `json:"epoch"`
}

type StoredData struct {
	Data
	Generation   uint32 `json:"generation"`
	TTL          uint32 `json:"ttl"`
	DiscardValue bool   `json:"-"`
	Gather       bool   `json:"-"`
}

type ComputeResponseV1 struct {
	StoredData
	Name string `json:"name"`
}

func NewComputeResponse(name string, key string, value []byte, epoch int64, generation uint32) ComputeResponseV1 {
	return ComputeResponseV1{
		Name: name,
		StoredData: StoredData{
			Generation: generation,
			Data: Data{
				Key:   key,
				Value: value,
				Epoch: epoch,
			},
		},
	}
}

func NewComputeResponseWithTTL(name string, key string, value []byte, epoch int64, generation, ttl uint32) ComputeResponseV1 {
	return ComputeResponseV1{
		Name: name,
		StoredData: StoredData{
			Generation: generation,
			TTL:        ttl,
			Data: Data{
				Key:   key,
				Value: value,
				Epoch: epoch,
			},
		},
	}
}
