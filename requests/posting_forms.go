package main

import (
	"net/http"
	"net/url"
	"log"
)

func main() {
	api := "http://localhost:5000/users/login"
	response, err := http.PostForm(api, url.Values{
		"email": { "joenyugoh@gmail.com" },
		"password": {  "****" },
	})
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // Clean after yourself :)

	log.Println(response.Header)
	log.Println(response.Status)
	log.Println(response.StatusCode)
	/*body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))*/



}
