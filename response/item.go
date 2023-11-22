package response

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Payload interface {
	Key() string
}

type Item struct {
	Type    string
	Payload Payload
}

var (
	ErrEmpty        = errors.New("emtpy response item")
	ErrMultipleKeys = errors.New("multiple keys in response item")
)

func (i *Item) UnmarshalJSON(data []byte) error {
	temp := map[string]json.RawMessage{}
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}
	switch len(temp) {
	case 0:
		return ErrEmpty
	case 1:
		for k := range temp {
			decoder, ok := itemRegistry[k]
			if !ok {
				return fmt.Errorf("no decoder for %q", k)
			}
			payload, err := decoder.decodeJSON(temp[k])
			if err != nil {
				return err
			}
			i.Payload = payload
		}
	default:
		return ErrMultipleKeys
	}
	return nil
}
