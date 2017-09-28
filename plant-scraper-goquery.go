package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const rhsBaseURL = "https://www.rhs.org.uk"

func getPage(url string) *goquery.Document {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Printf("Unable to build doc at url %v, err = %v", url, err)
	}

	return doc
}

func findPlantURLs(doc *goquery.Document) []string {
	var urls []string
	doc.Find(".posts-list-content .result-details .Plant-formated-Name").Each(func(index int, item *goquery.Selection) {
		relPath, _ := item.Attr("href")
		urls = append(urls, rhsBaseURL+relPath)
	})
	return urls
}

func findPlantName(doc *goquery.Document) string {
	name := doc.Find(".Plant-formated-Name").First().Text()
	return strings.TrimSpace(name)
}

func main() {
	bulbsDoc := getPage(rhsBaseURL + "/plants/bulbs")
	urls := findPlantURLs(bulbsDoc)
	for _, URL := range urls {
		plantDoc := getPage(URL)
		plantName := findPlantName(plantDoc)
		fmt.Printf("%v\n", plantName)
	}

}
