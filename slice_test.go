package sandboxrpc

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEmptySlice(t *testing.T) {
	var s Slice[int]
	raw, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, `[]`, string(raw))
}

func FuzzSlice(f *testing.F) {
	f.Add("test")
	f.Fuzz(func(t *testing.T, s string) {
		stringJSON, err := json.Marshal(s)
		require.NoError(t, err)
		sliceJSON, err := json.Marshal(Slice[string]{s})
		require.NoError(t, err)
		require.Equal(t, "["+string(stringJSON)+"]", string(sliceJSON))
	})
}
