package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/base64"
)

func main() {
	url := "https://api.splunknovadev-playground.com/v1/events"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`[{ "log": "This is my first Nova event", "source": "curl", "entity": "test_api" }]`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// hack
	data := []byte("<REDACTED>")
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
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}