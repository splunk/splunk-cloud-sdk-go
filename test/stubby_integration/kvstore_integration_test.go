package stubbyintegration


import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats("namespace1", "collection1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "5", result.Count)
	assert.Equal(t, "namespace1", result.Ns)
	assert.Equal(t, "1", result.Nindexes)
}

func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "healthy", result.Status)
}
