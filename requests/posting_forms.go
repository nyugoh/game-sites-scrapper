package main

import (
	"net/http"
	"net/url"
	"log"
	"io/ioutil"
)

func main() {
	api := "http://localhost:5000/users/login"
	response, err := http.PostForm(api, url.Values{
		"email": { "joenyugoh@gmail.com" },
		"password": {  "Shushume@20" },
	})
	if err != nil {
		log.Fatal(err)
	} else {
		//log.Println(response.Header)
		//log.Println(response.Cookies())
		//log.Println(response.StatusCode)

		goHome(response.Cookies())
	}
	defer response.Body.Close() // Clean after yourself :)
}

func goHome(authCookies []*http.Cookie)  {
	// Make a request to home page with the cookie
	homeRequest, err := http.NewRequest("GET", "http://localhost:5000", nil)

	if err != nil {
		log.Fatal(err)
	}

	for _, cookie := range authCookies {
		homeRequest.AddCookie(cookie)
	}

	log.Println(homeRequest.Header)
	log.Println(homeRequest.Cookies())
	// Make the request
	client := &http.Client{}
	jar := client.Jar
	url := &url.URL{}
	url.Scheme = "http"
	url.Host = "localhost:5000"
	url.Path = "/"

	jar.SetCookies(url.String(), authCookies)
	http.SetCookie()
	res, err := client.Do(homeRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	log.Println(res.Header)
	log.Println(res.StatusCode)
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
