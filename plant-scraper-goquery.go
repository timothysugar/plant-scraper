package main

import (
	"log"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func getDoc() *goquery.Document {
	doc, err := goquery.NewDocument("https://www.rhs.org.uk/plants/bulbs")
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func postScrape(doc *goquery.Document) []string {
	var urls []string
	doc.Find(".posts-list-content .result-details .Plant-formated-Name").Each(func(index int, item *goquery.Selection) {
		url, _ := item.Attr("href")
		urls = append(urls, url)
	})
	return urls
}

func main() {
	doc := getDoc()
	urls := postScrape(doc)
	fmt.Println(urls)
}
