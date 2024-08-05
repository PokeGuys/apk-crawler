package apkpure

import (
	"strings"

	apkcrawler "github.com/pokeguys/apk-crawler"
	"github.com/pokeguys/apk-crawler/sources/apkpure/apkpurehttp"
)

type Crawler struct {
	Config
	client apkpurehttp.Client
}

// NewApkPureCrawler returns a new Crawler for the ApkPure API.
func NewApkPureCrawler(cfg Config, client apkpurehttp.Client) (*Crawler, error) {
	return &Crawler{
		Config: cfg,
		client: client,
	}, nil
}

// Crawl sends a request to the API and returns the extracted apk information.
func (c *Crawler) Crawl(packageName, apkType string) ([]apkcrawler.Apk, error) {
	// 1. Send the get request to the API.
	// 2. Decode the response using the protobuf library.
	// 3. Iterate over the response and extract the apk information.
	// 4. Return the extracted apk information.
	apiResult, err := c.client.GetVersions(c.SDKVersion, c.Abis, packageName)
	if err != nil {
		return nil, err
	}

	// Check if the response is empty
	// All the dot are nullable
	apks := make([]apkcrawler.Apk, 0)
	if apiResult.Data == nil || apiResult.Data.Detail == nil || apiResult.Data.Detail.ApplicationVersion == nil {
		return apks, nil
	}

	// Transform the protobuf response into a list of Apk objects
	for _, app := range apiResult.Data.Detail.ApplicationVersion {
		// Skip the application if it's not the correct type
		if !strings.EqualFold(app.Result.Data.Download.Type, apkType) {
			continue
		}
		apks = append(apks, apkcrawler.Apk{
			Name:    app.Result.Data.Name,
			Package: app.Result.Data.Package,
			Version: app.Result.Data.Version,
			Size:    app.Result.Data.Download.Size,
			URL:     app.Result.Data.Download.Url,
			Hash:    app.Result.Data.Download.Sha1,
		})
	}
	// The response is already sorted in descending order
	return apks, nil
}
