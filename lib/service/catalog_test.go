package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_dataset(t *testing.T) {
	var splunkClient = NewSplunkdClient("",
		[2]string{"admin", "changeme"},
		"localhost:8882", "http", nil)

	result, err := splunkClient.CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, result.Name, "ds1")
	//assert.Equal(t, result.Kind, "")

}
