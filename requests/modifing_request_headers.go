package requests

import (
	"net/http"
	"time"
	"log"
	"io"
	"os"
	"fmt"
)

// To change headers is part of making your client specification :)
// I will add headers to impersonate Firefox browser

func main() {
	// Make the client
	client := &http.Client{
		Timeout: 30 * time.Second, // 30 seconds timeout
	}

	// Create a request which we will modify the headers
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add the headers to the request
	req.Header.Set("User-Agent", "Awesome web scrapper")

	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	if n, err := io.Copy(os.Stdout, res.Body) ; err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\n Bytes :: %d", n)
		fmt.Println(res.Header)
	}

}
