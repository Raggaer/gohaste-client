package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var serverURL = "https://hastebin.com"

type serverResponse struct {
	Message string `json:"message"`
	Key     string `json:"key"`
}

// Uploads the given data to a hastebin server
func Upload(data []byte) (*serverResponse, error) {
	buff := bytes.NewBuffer(data)
	resp, err := http.Post(serverURL+"/documents", "application/json", buff)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Retrieve JSON response from hastebin server
	var response serverResponse
	if err := json.Unmarshal(content, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
