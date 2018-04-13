package service

import (
	"github.com/splunk/ssc-client-go/lib/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getDataset(t *testing.T) {
	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
}

func Test_getDatasets(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

func Test_postDataset(t *testing.T) {
	result, err := getSplunkClient().CatalogService.PostDataset(
		getSplunkClient().CatalogService.CreateDataset("ds1", model.VIEW, []string{"string"}, "string"))
	assert.Empty(t, err)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
	assert.Equal(t, []string{"string"}, result.Rules)
}

func Test_deleteDataset(t *testing.T) {
	err := getSplunkClient().CatalogService.DeleteDataset("ds1")
	assert.Empty(t, err)
}
