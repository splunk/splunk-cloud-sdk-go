package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/base64"
)

func main() {
	//url := "https://api.splunknovadev-playground.com/v1/events"
	url := "https://api.splunknovadev-playground.com/search/v1/jobs"
	fmt.Println("URL:>", url)

	//var jsonStr = []byte(`[{ "log": "This is my first Nova event", "source": "curl", "entity": "test_api" }]`)

	var jsonStr = []byte(`[{ "query": "search index = _internal"}]`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// hack
	data := []byte("<REDACTED>:<REDACTED>")
	str := base64.StdEncoding.EncodeToString(data)

	req.Header.Set("Authorization", "Basic " + str)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	// TODO(dan): this seems like a better way, not sure why we went with def because this is what ReadAll does
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}