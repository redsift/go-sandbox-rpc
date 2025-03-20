package sandboxrpc

import "encoding/json"

// Slice is a workaround for empty slices and nil slices being exactly the same, except when
// marshaling to JSON.
type Slice[T any] []T

func (s Slice[T]) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte(`[]`), nil
	}
	return json.Marshal([]T(s))
}
