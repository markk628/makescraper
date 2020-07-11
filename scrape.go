package main

import (
	"fmt"
	"log"
	"github.com/gocolly/colly"
)

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
                link := e.Attr("href")

		// Print link
                fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[#t3_hf1web > div._1poyrkZ7g36PawDueRza-J._11R7M_VOgKO1RJyRSRErT3 > div._2FCtq-QzlfuN-SwVMUZMM3._3wiKjmhpIpoTE2r5KCm2o6.t3_hf1web > div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE > a]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	
	c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.reddit.com/")
}
