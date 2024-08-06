package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"

	"github.com/pokeguys/apk-crawler/cmd/crawler/config"
	"github.com/pokeguys/apk-crawler/sources"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("Usage: %s <source> [flags]\n", os.Args[0])
		os.Exit(1)
	}

	client := resty.New()
	cfg := config.NewConfig(args)

	// Setup the http client for the crawler
	crawler, err := sources.NewCrawlerStrategy(cfg.Source, client)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Crawl the package
	apk, err := crawler.Crawl(cfg.Package, cfg.ApkType)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print the results in JSON format
	if cfg.PrintJSON {
		enc := json.NewEncoder(os.Stdout)
		buf := new(bytes.Buffer)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(apk); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(string(buf.Bytes()))
		return
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
