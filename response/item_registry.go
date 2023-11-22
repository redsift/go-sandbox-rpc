package response

import (
	"encoding/json"
	"fmt"
)

type registryItem struct {
	decodeJSON func(json.RawMessage) (Payload, error)
}

var itemRegistry = map[string]registryItem{}

func RegisterItem[Type Payload](i Type) error {
	if _, ok := itemRegistry[i.Key()]; ok {
		return fmt.Errorf("multiple registrations for payload key %q", i.Key())
	}

	itemRegistry[i.Key()] = registryItem{
		decodeJSON: func(raw json.RawMessage) (Payload, error) {
			var item Type
			err := json.Unmarshal(raw, &item)
			return item, err
		},
	}

	return nil
}
