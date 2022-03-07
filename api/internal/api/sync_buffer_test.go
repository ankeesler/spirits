package api

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyncBuffer(t *testing.T) {
	sb := syncBuffer{}
	require.Equal(t, "", sb.read())

	n, err := sb.Write([]byte("tuna"))
	require.Equal(t, 4, n)
	require.NoError(t, err)

	require.Equal(t, "tuna", sb.read())
	require.Equal(t, "", sb.read())
	require.Equal(t, "", sb.read())

	n, err = sb.Write([]byte("fish"))
	require.Equal(t, 4, n)
	require.NoError(t, err)

	require.Equal(t, "fish", sb.read())
	require.Equal(t, "", sb.read())
	require.Equal(t, "", sb.read())

	n, err = sb.Write([]byte("marlin"))
	require.Equal(t, 6, n)
	require.NoError(t, err)
	n, err = sb.Write([]byte("squid"))
	require.Equal(t, 5, n)
	require.NoError(t, err)

	require.Equal(t, "marlinsquid", sb.read())
	require.Equal(t, "", sb.read())
	require.Equal(t, "", sb.read())
}
