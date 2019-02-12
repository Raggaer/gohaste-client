package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

var retrieveRaw bool

func main() {
	// Parse flags
	flag.StringVar(&serverURL, "server", "https://hastebin.com", "Hastebin server URL")
	flag.BoolVar(&retrieveRaw, "raw", false, "Retrieve raw URL")
	flag.Parse()

	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var content []byte

	// Determine if input comes from stdin
	// Or if input should be an argument file
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		// Retrieve file to read from command line arguments
		if len(flag.Args()) <= 0 {
			fmt.Println("Specify a file to read or pipe data to this application")
			return
		}
		fileName := flag.Args()[0]
		content, err = ioutil.ReadFile(fileName)
	} else {
		// Read everything from stdin
		content, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		panic(err)
	}

	resp, err := upload(content)
	if err != nil {
		fmt.Println("Unable to send data to " + serverURL)
		fmt.Println(err.Error())
		return
	}

	// Check for errors
	if resp.Message != "" {
		fmt.Println("Unable to send data to " + serverURL)
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Document uploaded to " + serverURL)
	finalURL := serverURL + "/" + resp.Key
	if retrieveRaw {
		finalURL = serverURL + "/raw/" + resp.Key
	}

	// Try to write result to clipboard if possible
	if err := clipboard.WriteAll(finalURL); err != nil {
		fmt.Println("- " + finalURL)
		fmt.Println("Unable to copy link to clipboard")
		fmt.Println(err.Error())
	} else {
		fmt.Println("- " + finalURL + " (copied to clipboard)")
	}
}
