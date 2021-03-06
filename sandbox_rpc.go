//
//  sandbox_rpc.go
//  Sandbox RPC model
//
//  Copyright (c) 2017 RedSift Limited. All rights reserved.
//
//
package sandboxrpc

type DataQuantum interface {
	Stored() []*StoredData
}

// ComputeRequest contains the parameters to invoke the node implementation function.
type ComputeRequest struct {
	In    *StoredDataQuantum `json:"in"`
	Query []string           `json:"query,omitempty"`
	With  *StoredDataQuantum `json:"with,omitempty"`
	Get   []GetDataQuantum   `json:"get,omitempty"`
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

type GetDataQuantum struct {
	Bucket string        `json:"bucket"`
	Key    string        `json:"key"`
	Data   []*StoredData `json:"data"`
}

func (d *GetDataQuantum) Stored() []*StoredData {
	return d.Data
}

func (d *StoredDataQuantum) Stored() []*StoredData {
	return d.Data
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

type ComputeResponse struct {
	StoredData
	Name string `json:"name"`
}

func NewComputeResponse(name string, key string, value []byte, epoch int64, generation uint32) ComputeResponse {
	return ComputeResponse{
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

func NewComputeResponseWithTTL(name string, key string, value []byte, epoch int64, generation, ttl uint32) ComputeResponse {
	return ComputeResponse{
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
