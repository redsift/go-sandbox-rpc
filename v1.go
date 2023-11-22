// sandbox_rpc.go
// Sandbox RPC model
//
// Copyright (c) 2017 RedSift Limited. All rights reserved.
package sandboxrpc

import legacy "github.com/redsift/go-sandbox-rpc"

type ComputeRequestV1 = legacy.ComputeRequest
type ResponseV1 = legacy.Response
type StoredDataQuantum = legacy.StoredDataQuantum
type Data = legacy.Data
type StoredData = legacy.StoredData
type ComputeResponseV1 = legacy.ComputeResponse

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
