package main

import (
	"net/http"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	// Close the stream
	defer req.Body.Close()

	n, err := io.Copy(os.Stdout, req.Body)
	if err != nil {
		log.Fatal(err)
	}
	 fmt.Printf("Fetched %d bytes", n)

}
