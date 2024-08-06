package sources

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	apkcrawler "github.com/pokeguys/apk-crawler"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
	"github.com/pokeguys/apk-crawler/sources/apkpure/apkpurehttp"
)

// Strategy is the interface for a source strategy.
type CrawlerStrategy interface {
	Crawl(packageName, apkType string) ([]apkcrawler.Apk, error)
}

// SourceConfig is the interface for a source configuration.
type SourceConfig interface {
	Name() string
	Validate() string
}

// NewStrategy returns a new strategy for the given source.
func NewCrawlerStrategy(config interface{}, client *resty.Client) (CrawlerStrategy, error) {
	switch cfg := config.(type) {
	case apkpure.Config:
		return apkpure.NewApkPureCrawler(cfg, apkpurehttp.NewClient(client))
	default:
		return nil, fmt.Errorf("unknown config type: %T", cfg)
	}
}

func NewSourceConfig(name string) (SourceConfig, error) {
	switch name {
	case apkpure.ConfigName():
		return apkpure.Config{}, nil
	default:
		return nil, fmt.Errorf("unknown source: %s", name)
	}
}
