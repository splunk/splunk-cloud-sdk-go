package stubbyintegration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Stubby test for GetCollectionStats() kvstore service endpoint
func TestGetCollectionStats(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetCollectionStats("namespace1", "collection1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, int64(5), result.Count)
	assert.Equal(t, "namespace1", result.Ns)
	assert.Equal(t, int64(1), result.Nindexes)
}

// Stubby test for ping/GetServiceHealthStatus() kvstore service endpoint
func TestGetServiceHealthStatus(t *testing.T) {
	result, err := getClient(t).KVStoreService.GetServiceHealthStatus()
	assert.Empty(t, err)
	assert.NotEmpty(t, result)

	assert.Equal(t, "healthy", result.Status)
}
