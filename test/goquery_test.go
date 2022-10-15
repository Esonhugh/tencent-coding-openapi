package test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestGetOpenAPI(t *testing.T) {
	// Request the HTML page.
	res, err := http.Get("https://coding.net/help/openapi")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("Openapi.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Find the review items
	doc.Find("#__NEXT_DATA__").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		text := s.Text()
		fmt.Printf("Review %d: %s\n", i, text)
		file.WriteString(text)
	})
}
