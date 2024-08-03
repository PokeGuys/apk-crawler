package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"

	"github.com/pokeguys/apk-crawler/cmd/crawler/config"
	"github.com/pokeguys/apk-crawler/sources"
)

func main() {
	cfg := config.NewConfig()

	// Setup the http client for the crawler
	client := resty.New()
	crawler := sources.NewCrawlerStrategy(cfg.Source, client)

	// Crawl the package
	apk, err := crawler.Crawl(cfg.Package)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print the results
	// If the user pass showAll flag, print all the results
	if cfg.ShowAll {
		for _, a := range apk {
			fmt.Println(a.URL)
		}
	} else {
		fmt.Println(apk[0].URL)
	}
}
