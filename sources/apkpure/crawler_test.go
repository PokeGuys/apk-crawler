package apkpure_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"google.golang.org/protobuf/proto"

	pb "github.com/pokeguys/apk-crawler/proto"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
)

func mockApkPureResponse(packageName, apkType string, count int) *pb.ApkPureResponse {
	apiResult := pb.ApkPureResponse{
		Data: &pb.ApkPureResponseData{
			Detail: &pb.ApkPureResponseDetail{},
		},
	}
	for i := 0; i < count; i++ {
		apiResult.Data.Detail.ApplicationVersion = append(
			apiResult.Data.Detail.ApplicationVersion,
			&pb.ApkPureApplicationVersion{
				Result: &pb.ApkPureApplicationSearchResult{
					Data: &pb.ApkPureApplicationVersionData{
						Name:         "Example App",
						DisplayName:  "Example App",
						Package:      packageName,
						MinorVersion: "1.0",
						Version:      "1.0.0",
						Hash:         "1234567890",
						Description:  "This is an example app",
						PatchNotes:   "This is an example app",
						Status:       "Published",
						Developer:    "Example Developer",
						Download: &pb.ApkPureApplicationDownload{
							Name:       "Example App",
							Sha1:       "1234567890",
							Size:       1024,
							TorrentUrl: "https://example.com/example.apk",
							TrackerUrl: "https://example.com/example.apk",
							Type:       apkType,
							Url:        "https://example.com/example.apk",
						},
					},
				},
			},
		)
	}
	return &apiResult
}

func TestCrawler_Crawl(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	crawler := apkpure.NewApkPureCrawler(client)

	// Marshal a protobuf response
	packageName := "com.example.app"
	apiResponse := mockApkPureResponse(packageName, "APK", 5)
	response, err := proto.Marshal(apiResponse)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Set the mock response
	httpmock.RegisterResponder("GET", crawler.URL(packageName),
		httpmock.NewStringResponder(200, string(response)))

	// Test the crawler with a valid package name
	apks, err := crawler.Crawl("com.example.app")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks) != 5 {
		t.Errorf("Expected 1 apk, got %d", len(apks))
	}

	// Verify the first apk
	if apks[0].Package != packageName {
		t.Errorf("Expected package to be '%s', got %s", packageName, apks[0].Package)
	}
}

func TestCrawler_CrawlAllXAPK(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	crawler := apkpure.NewApkPureCrawler(client)

	// Marshal a protobuf response
	packageName := "com.example.app"
	response, err := proto.Marshal(mockApkPureResponse(packageName, "XAPK", 2))
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Set the mock response
	httpmock.RegisterResponder("GET", crawler.URL(packageName),
		httpmock.NewStringResponder(200, string(response)))

	// Test the crawler with a valid package name
	apks, err := crawler.Crawl("com.example.app")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks) != 0 {
		t.Errorf("Expected 0 apk, got %d (filtered out XAPK)", len(apks))
	}
}

func TestCrawler_CrawlEmptyResponse(t *testing.T) {
	client := resty.New()
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	crawler := apkpure.NewApkPureCrawler(client)
	var apiResult pb.ApkPureResponse
	response, err := proto.Marshal(&apiResult)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Set the mock response
	httpmock.RegisterResponder("GET", crawler.URL("com.example.app"),
		httpmock.NewStringResponder(200, string(response)))

	// Test the crawler with a valid package name
	apks, err := crawler.Crawl("com.example.app")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks) != 0 {
		t.Errorf("Expected 0 apk, got %d", len(apks))
	}
}
