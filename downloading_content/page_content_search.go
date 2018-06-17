package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"strings"
	"os"
)

/*
 We are going to get a page via request,  http.Get
convert it into a string that can be searched using ioutil.ReadAll
search for anything in the string i.e page response using either strings li or req exps
 */
func main() {
	req, err := http.Get("http://magnumdigitalke.com")
	if err != nil {
		log.Fatal(err)
	}

	defer req.Body.Close()

	// Convert the response to string
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	pageContent := string(body)

	// Let's search for the index of <title>, this is where we will the begining of the title
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		log.Fatal("No title on this page")
		os.Exit(0)
	}
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		log.Fatal("No closing title")
	}

	// My name
	nameIndex := strings.Index(pageContent, "Nyugoh")
	if nameIndex == -1 {
		log.Fatal("Not found")
	} else {
		fmt.Println(pageContent[nameIndex-5:nameIndex+8])
	}

	// Using start end we can get the title content
	title := pageContent[titleStartIndex+7:titleEndIndex]

	fmt.Println(title)
}
