package main

import (
	"net/http"
	"time"
	"log"
	"os"
	"io"
	"fmt"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close() // Close stream later

	// Make the output directory and file to save content
	os.Mkdir("output", 0777) // will be created in the project root
	outputFile, err := os.Create("output/google.html")
	if err != nil {
		log.Fatal(err)
	}

	// Write to response to file
	_, error := io.Copy(outputFile, res.Body)
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println("Done writing to file.")
	}

}
