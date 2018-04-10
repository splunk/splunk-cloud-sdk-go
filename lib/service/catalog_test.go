package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/splunk/ssc-client-go/lib/model"
	"time"
)

func getSplunkClient() *Client {

	return NewClient("",
		[2]string{"admin", "changeme"},
		"localhost:8882", "http", time.Second*5, true)
}

func Test_getDataset(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDataset("ds1")
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
}

func Test_getDatasets(t *testing.T) {

	result, err := getSplunkClient().CatalogService.GetDatasets()
	assert.Empty(t, err)
	assert.Equal(t, 2, len(result))
}

func Test_postDataset(t *testing.T) {

	//dataset := model.Dataset_post{"ds1", model.VIEW, []string{"string"}, "string"}
	result, err := getSplunkClient().CatalogService.PostDataset(
		getSplunkClient().CatalogService.CreateDataset("ds1", model.VIEW, []string{"string"}, "string"))
	assert.Empty(t, err)
	assert.NotEmpty(t, result.Id)
	assert.Equal(t, "ds1", result.Name)
	assert.Equal(t, model.VIEW, result.Kind)
	assert.Equal(t, []string{"string"}, result.Rules)
}

func Test_deleteDataset(t *testing.T) {
	err := getSplunkClient().CatalogService.DeleteDataset("ds1")
	assert.Empty(t, err)
}
