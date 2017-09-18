//
//  sandbox_rpc.go
//  Sandbox RPC model
//
//  Copyright (c) 2017 RedSift Limited. All rights reserved.
//
//
package sandboxrpc

// ComputeRequest contains the parameters to invoke the node implementation function.
type ComputeRequest struct {
	In          *StoredDataQuantum  `json:"in"`
	Aggregation []string            `json:"query,omitempty"`
	With        *StoredDataQuantum  `json:"with,omitempty"`
	Lookup      []LookupDataQuantum `json:"lookup,omitempty"`
}

// Response is returned by the sandbox.
type Response struct {
	Out   []*ComputeResponse `json:"out,omitempty"`
	Error interface{}        `json:"error,omitempty"`
	Stats interface{}        `json:"stats,omitempty"` // timings and resource usage to be booked against this sift TODO: nail down
}

// A single unit of operation
type StoredDataQuantum struct {
	Bucket string        `json:"bucket"`
	Data   []*StoredData `json:"data"`
}

type LookupDataQuantum struct {
	Bucket string      `json:"bucket"`
	Data   *StoredData `json:"data"`
}

type Data struct {
	Key   string `json:"key"`
	Value []byte `json:"value"`
	Epoch int64  `json:"epoch"`
}

type StoredData struct {
	Data
	Generation   uint32 `json:"generation"`
	DiscardValue bool   `json:"-"`
	Gather       bool   `json:"-"`
}

type ComputeResponse struct {
	StoredData
	Name string `json:"name"`
}

func NewComputeResponse(n string, k string, v []byte, e int64, g uint32) ComputeResponse {
	return ComputeResponse{
		Name: n,
		StoredData: StoredData{
			Generation: g,
			Data: Data{
				Key:   k,
				Value: v,
				Epoch: e,
			},
		},
	}
}
