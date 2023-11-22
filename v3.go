package sandboxrpc

import (
	"encoding/json"
	"fmt"

	"github.com/redsift/go-sandbox-rpc/v3/response"
	"github.com/redsift/go-sandbox-rpc/v3/sandboxrpc"
)

type ComputeRequestV3 struct {
	Secrets map[string]any        `json:"secrets" msgpack:"secrets"`
	Stores  map[string]StoreData  `json:"stores" msgpack:"stores"`
	Streams map[string]StreamSpec `json:"streams" msgpack:"streams"`
}

type StreamSpec = sandboxrpc.StreamSpec

func (cr *ComputeRequestV3) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Secrets map[string]any
		Stores  map[string]map[string]json.RawMessage
	}{}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	cr.Secrets = tmp.Secrets
	cr.Stores = make(map[string]StoreData, len(tmp.Stores))

	for store, tmpStore := range tmp.Stores {
		if tmpData, ok := tmpStore["data"]; ok {
			isd := new(InlineStoreData)
			if err := json.Unmarshal(tmpData, isd); err != nil {
				return fmt.Errorf("cannot unmarshal store %q: %w", store, err)
			}
			cr.Stores[store] = isd
		} else {
			return fmt.Errorf("cannot unmarshal store %q: unknown store type: %v", store, tmpStore)
		}
	}
	return nil
}

type ComputeResponseV3 = response.ComputeResponse
