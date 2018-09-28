package link

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapped_Bytes(t *testing.T) {
	tt := getTestMaps()

	for _, v := range tt {

		buf := v.m.Bytes()

		mapped, err := BytesToMapped(buf)
		require.NoError(t, err)

		require.Equal(t, v.m, mapped)
	}
}
