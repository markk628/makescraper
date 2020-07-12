package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Post struct {
	Title string
	PostDate string
	Votes string
	Comments string
}

type Posts struct {
	Posts string
}

func jsonFile(filename string, e Posts) {
	jsonPost, err := json.Marshal(e)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("output.json", jsonPost, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	c := colly.NewCollector()
	var posts [] Posts
	
	c.OnHTML("div._2SdHzo12ISmrC8H86TgSCp._3wqmjmv3tb_k-PROt7qFZe", func(e *colly.HTMLElement) {

		post := Posts{Posts: e.Text}

		j, _ := json.Marshal(post)
		_ = ioutil.WriteFile("output.json", j, os.ModePerm)

		for _, post := range posts {
			jsonFile("output.json", post)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error")
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.reddit.com/")

	// Instantiate default collector
	// c := colly.NewCollector()

	// // On every a element which has href attribute call callback
	// c.OnHTML("div._2SdHzo12ISmrC8H86TgSCp._3wqmjmv3tb_k-PROt7qFZe", func(e *colly.HTMLElement) {
	// 	// Print link
    //     fmt.Printf("Title: %s\n", e.Text)
	// })

	// c.OnHTML("a._3jOxDPIQ0KaOWpzvSQo-1s", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Posted: %s\n", e.Text)
	// })

	// c.OnHTML("div._1rZYMD_4xY3gRcSS3p8ODO", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Votes: %s\n", e.Text)
	// })

	// c.OnHTML("a._1UoeAeSRhOKSNdY_h3iS1O._1Hw7tY9pMr-T1F4P1C-xNU._2qww3J5KKzsD7e5DO0BvvU", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Comments: %s\n", e.Text)
	// })

	// // Before making a request print "Visiting ..."
	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	// c.OnError(func(_ *colly.Response, err error) {
	// 	log.Println("Something went wrong:", err)
	// })
	
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited", r.Request.URL)
	// })

	// c.OnHTML("a[#t3_hf1web > div._1poyrkZ7g36PawDueRza-J._11R7M_VOgKO1RJyRSRErT3 > div._2FCtq-QzlfuN-SwVMUZMM3._3wiKjmhpIpoTE2r5KCm2o6.t3_hf1web > div.y8HYJ-y_lTUHkQIc1mdCq._2INHSNB8V5eaWp4P0rY_mE > a]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	
	// c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	// 	fmt.Println("First column of a table row:", e.Text)
	// })
	
	// c.OnXML("//h1", func(e *colly.XMLElement) {
	// 	fmt.Println(e.Text)
	// })
	
	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("Finished", r.Request.URL)
	// })

	// // Start scraping on https://hackerspaces.org
	// c.Visit("https://www.reddit.com/")
}
