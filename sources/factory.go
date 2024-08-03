package sources

import (
	"github.com/go-resty/resty/v2"

	apkcrawler "github.com/pokeguys/apk-crawler"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
)

// Strategy is the interface for a source strategy.
type CrawlerStrategy interface {
	Crawl(packageName string) ([]apkcrawler.Apk, error)
}

// NewStrategy returns a new strategy for the given source.
func NewCrawlerStrategy(name string, client *resty.Client) CrawlerStrategy {
	// Convert the source name to a source enum
	source := GetSourceByName(name)
	switch source {
	case ApkPure:
		return apkpure.NewApkPureCrawler(client)
	default:
		panic("unknown source")
	}
}
