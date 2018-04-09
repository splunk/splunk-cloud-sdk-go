package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getSplunkClient() *Client {

	return NewSplunkdClient("",
		[2]string{"admin", "changeme"},
		"localhost:8882", "http", nil)
}

func Test_getDataset(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, result.Name, "ds1")
	assert.Equal(t, result.Kind, "VIEW")
}

func Test_getDatasets(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, len(result), 2)
}
