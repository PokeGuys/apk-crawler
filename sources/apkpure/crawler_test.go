package apkpure_test

import (
	"testing"

	mocks "github.com/pokeguys/apk-crawler/mocks/github.com/pokeguys/apk-crawler/sources/apkpure/apkpurehttp"
	pb "github.com/pokeguys/apk-crawler/proto"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
	"github.com/pokeguys/apk-crawler/sources/apkpure/apkpurehttp"
)

func TestCrawler_Crawl(t *testing.T) {
	client := &mocks.MockClient{}
	packageName := "com.example.app"
	sdkVersion := "29"
	abis := "x86,armeabi-v7a,arm64-v8a,x86_64"
	client.On("GetVersions", sdkVersion, abis, packageName).Return(apkpurehttp.MockApkPureResponse(packageName, "APK", 5), nil)

	// Test the crawler with a valid package name
	crawler, err := apkpure.NewApkPureCrawler(apkpure.Config{
		SDKVersion: sdkVersion,
		Abis:       abis,
	}, client)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	apks, err := crawler.Crawl(packageName, "APK")
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
	client := &mocks.MockClient{}
	packageName := "com.example.app"
	sdkVersion := "29"
	abis := "x86,armeabi-v7a,arm64-v8a,x86_64"
	client.On("GetVersions", sdkVersion, abis, packageName).Return(apkpurehttp.MockApkPureResponse(packageName, "XAPK", 2), nil)

	// Test the crawler with a valid package name
	crawler, err := apkpure.NewApkPureCrawler(apkpure.Config{
		SDKVersion: sdkVersion,
		Abis:       abis,
	}, client)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	apks, err := crawler.Crawl(packageName, "APK")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks) != 0 {
		t.Errorf("Expected 0 apk, got %d (filtered out XAPK)", len(apks))
	}
}

func TestCrawler_CrawlEmptyResponse(t *testing.T) {
	client := &mocks.MockClient{}
	packageName := "com.example.app"
	sdkVersion := "29"
	abis := "x86,armeabi-v7a,arm64-v8a,x86_64"
	var apiResult pb.ApkPureResponse
	client.On("GetVersions", sdkVersion, abis, packageName).Return(&apiResult, nil)
	// Test the crawler with a valid package name

	crawler, err := apkpure.NewApkPureCrawler(apkpure.Config{
		SDKVersion: sdkVersion,
		Abis:       abis,
	}, client)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	apks, err := crawler.Crawl(packageName, "APK")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks) != 0 {
		t.Errorf("Expected 0 apk, got %d", len(apks))
	}
}
