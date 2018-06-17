package main

import (
	"time"
	"log"
	"os"
	"io"
	"fmt"
	"net/http"
)

func main() {
	url := "https://www.sportpesa.co.ke/sportgames?sportId=1"
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	backupFile, err := os.Create("games.html")
	if err != nil {
		log.Fatal(err)
	}

	n, err := io.Copy(backupFile, req.Body)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Saved to file %s bytes.", n)
	}
}
