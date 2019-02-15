package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUpload(t *testing.T) {
	// Upload some random text
	content, err := Upload([]byte("Hello World!"))
	if err != nil {
		t.Fatalf("Unable to send data to hastebin server %s", err)
		return
	}

	// Check if response contains an error
	if content.Message != "" {
		t.Fatalf("Unable to send data to hastebin server %s", content.Message)
		return
	}

	// Retrieve content from our raw pastebin
	resp, err := http.Get(serverURL + "/raw/" + content.Key)
	if err != nil {
		t.Fatalf("Unable to retrieve data from hastebin server %s", err)
		return
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unable to read data to hastebin server %s", err)
		return
	}

	// Check if paste content is the expected one
	if string(bodyContent) != "Hello World!" {
		t.Fatalf("Unexpected paste content, got '%s' expected '%s'", string(bodyContent), "Hello World!")
		return
	}
}
