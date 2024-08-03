package apkpure

import (
	"fmt"
	"sort"
	"strings"

	"github.com/go-resty/resty/v2"
	"google.golang.org/protobuf/proto"

	apkcrawler "github.com/pokeguys/apk-crawler"
	pb "github.com/pokeguys/apk-crawler/proto"
)

type Crawler struct {
	client *resty.Client
}

// NewApkPureCrawler returns a new Crawler for the ApkPure API.
func NewApkPureCrawler(client *resty.Client) *Crawler {
	return &Crawler{
		client: client,
	}
}

// Crawl sends a request to the API and returns the extracted apk information.
func (c *Crawler) Crawl(packageName string) ([]apkcrawler.Apk, error) {
	// 1. Send the get request to the API.
	// 2. Decode the response using the protobuf library.
	// 3. Iterate over the response and extract the apk information.
	// 4. Return the extracted apk information.
	resp, err := c.client.R().SetHeaders(c.header()).Get(c.url(packageName))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.RawBody().Close()

	var apiResult pb.ApkPureResponse
	err = proto.Unmarshal(resp.Body(), &apiResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Transform the protobuf response into a list of Apk objects
	apks := make([]apkcrawler.Apk, 0, len(apiResult.Data.Detail.ApplicationVersion))
	for _, app := range apiResult.Data.Detail.ApplicationVersion {
		// Skip the application if it's not APK
		if strings.ToUpper(app.Result.Data.Download.Type) != "APK" {
			continue
		}
		apks = append(apks, apkcrawler.Apk{
			Name:      app.Result.Data.Name,
			Version:   app.Result.Data.Version,
			Size:      app.Result.Data.Download.Size,
			URL:       app.Result.Data.Download.Url,
			Hash:      app.Result.Data.Download.Hash,
			CreatedAt: app.Result.Data.Download.CreatedAt,
		})
	}

	// Sort the apks by timestamp (newest first)
	sort.Slice(apks, func(i, j int) bool {
		return apks[i].CreatedAt < apks[j].CreatedAt
	})

	return apks, nil
}

func (c *Crawler) url(packageName string) string {
	return fmt.Sprintf(APIURL, packageName)
}

func (c *Crawler) header() map[string]string {
	return map[string]string{
		"x-cv":   "3172501",
		"x-sv":   "29",
		"x-abis": "x86,armeabi-v7a,arm64-v8a,x86_64",
		"x-gp":   "1",
	}
}
