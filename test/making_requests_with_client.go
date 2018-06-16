// To have more control over how requests are made by Go, making a custom http client is the way to go
package main

import (
	"net/http"
	"time"
	"log"
	"io"
	"os"
	"fmt"
)

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second, // Time out after 30 seconds
	}

	req, err := client.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer req.Body.Close()

	if n, err := io.Copy(os.Stdout, req.Body); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(n)
	}

}
