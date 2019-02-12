package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var content []byte

	// Determine if input comes from stdin
	// Or if input should be an argument file
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		// Retrieve file to read from command line arguments
		if len(os.Args) <= 1 {
			fmt.Println("Specify a file to read or pipe data to this application")
			return
		}
		fileName := os.Args[1]
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

	// Try to write result to clipboard if possible
	if err := clipboard.WriteAll(serverURL + "/" + resp.Key); err != nil {
		fmt.Println("- " + serverURL + "/" + resp.Key)
		fmt.Println("Unable to copy link to clipboard")
		fmt.Println(err.Error())
	} else {
		fmt.Println("- " + serverURL + "/" + resp.Key + "(copied to clipboard)")
	}
}
