package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	// Check if you go anything
	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	ul := doc.Find("ul")
	ul.Find("li").Each(func(idx int, item *goquery.Selection) {
		linkTag := item.Find("a")
		title := linkTag.Text()
		link, _ := linkTag.Attr("href")
		fmt.Printf("Post #%d: %s - %s\n", idx, title, link)
	})
}
