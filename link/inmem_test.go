package link

import (
	"testing"

	"github.com/adelowo/mapped/utils"
	"github.com/stretchr/testify/require"
)

var _ Repository = (*Memory)(nil)

func TestMemory_Create(t *testing.T) {

	m := NewMemoryRepository()

	tt := getTestMaps()

	for _, v := range tt {
		require.NoError(t, m.Create(v.m))
	}

	// Add them again, there should still be no errors
	for _, v := range tt {
		require.NoError(t, m.Create(v.m))
	}

	require.Equal(t, int64(len(tt)), m.len())
}

func TestMemory_Find(t *testing.T) {
	m := NewMemoryRepository()

	for _, v := range getTestMaps() {
		_, err := m.Find(v.m.Original)
		require.True(t, IsLinkNotFoundError(err))

		require.NoError(t, m.Create(v.m))
	}
}

func getTestMaps() []struct{ m Mapped } {
	return []struct {
		m Mapped
	}{
		{Mapped{
			Original:  "https://lanre.wtf",
			TO:        "https://lanre.wtf",
			CreatedAt: utils.Now(),
		}},
		{Mapped{
			Original:  "https://lanre.date",
			TO:        "https://lanre.wtf",
			CreatedAt: utils.Now(),
		}},
		{Mapped{
			Original:  "https://lanre.tech",
			TO:        "https://lanre.wtf",
			CreatedAt: utils.Now(),
		}},
		{Mapped{
			Original:  "https://lanre.fun",
			TO:        "https://lanre.wtf",
			CreatedAt: utils.Now(),
		}},
	}

}
