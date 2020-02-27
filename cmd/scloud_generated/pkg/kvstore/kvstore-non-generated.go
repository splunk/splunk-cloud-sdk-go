package kvstore

import (
	"encoding/json"
	"io/ioutil"

	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	model "github.com/splunk/splunk-cloud-sdk-go/services/kvstore"
)



// PostEvents Sends events.
func InsertRecordOverride(collection string, bodyfile string) (*model.Key, error) {
	var resp *model.Key

	client, err := auth.GetClient()
	if err != nil {
		return nil, err
	}

	byets, err := ioutil.ReadFile(bodyfile)
	if err != nil {
		return nil, err
	}

	var data  map[string]interface{}
	err = json.Unmarshal(byets, &data)
	if err != nil {
		return nil, err
	}


	resp, err = client.KVStoreService.InsertRecord(collection,data)

	return resp, nil
}
