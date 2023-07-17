// Copyright (c) 2023 RedSift Limited. All rights reserved.
package sandboxrpc

type DataQuantum interface {
	Stored() []*StoredData
}

// ComputeRequest contains the parameters to invoke the node implementation function.
type ComputeRequest struct {
	In    *StoredDataQuantum `json:"in" msgpack:"in"`
	Query []string           `json:"query,omitempty" msgpack:"query"`
	With  *StoredDataQuantum `json:"with,omitempty" msgpack:"with"`
	Get   []GetDataQuantum   `json:"get,omitempty" msgpack:"get"`
}

// Response is returned by the sandbox.
type Response struct {
	Out   []*ComputeResponse `json:"out,omitempty" msgpack:"out"`
	Error interface{}        `json:"error,omitempty" msgpack:"error"`
	Stats interface{}        `json:"stats,omitempty" msgpack:"stats"` // timings and resource usage to be booked against this sift TODO: nail down
}

// A single unit of operation
type StoredDataQuantum struct {
	Bucket string        `json:"bucket" msgpack:"bucket"`
	Data   []*StoredData `json:"data" msgpack:"data"`
}

type GetDataQuantum struct {
	Bucket string        `json:"bucket" msgpack:"bucket"`
	Key    string        `json:"key" msgpack:"key"`
	Data   []*StoredData `json:"data" msgpack:"data"`
}

func (d *GetDataQuantum) Stored() []*StoredData {
	return d.Data
}

func (d *StoredDataQuantum) Stored() []*StoredData {
	return d.Data
}

type Data struct {
	Key   string `json:"key" msgpack:"key"`
	Value []byte `json:"value" msgpack:"value"`
	Epoch int64  `json:"epoch" msgpack:"epoch"`
}

type StoredData struct {
	Data
	Generation   uint32 `json:"generation" msgpack:"generation"`
	TTL          uint32 `json:"ttl" msgpack:"ttl"`
	DiscardValue bool   `json:"-"`
	Gather       bool   `json:"-"`
}

type ComputeResponse struct {
	StoredData
	Name string `json:"name" msgpack:"name"`
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
