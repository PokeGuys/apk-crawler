package apkpurehttp_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"google.golang.org/protobuf/proto"

	"github.com/pokeguys/apk-crawler/sources/apkpure/apkpurehttp"
)

func TestClient_GetVersions(t *testing.T) {
	packageName := "com.example.app"
	apkType := "APK"
	// Set the mock response
	data, err := proto.Marshal(apkpurehttp.MockApkPureResponse(packageName, apkType, 5))
	if err != nil {
		t.Errorf("Expected no error when marshaling the response, got %v", err)
	}
	httpmock.RegisterResponder("GET", apkpurehttp.GetSearchURL(packageName),
		httpmock.NewStringResponder(200, string(data)))

	restyClient := resty.New()
	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()

	// Test the client with a valid package name
	client := apkpurehttp.NewClient(restyClient)
	apks, err := client.GetVersions("29", "x86,armeabi-v7a,arm64-v8a,x86_64", packageName)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Verify the response
	if len(apks.Data.Detail.ApplicationVersion) != 5 {
		t.Errorf("Expected 5 apk, got %d", len(apks.Data.Detail.ApplicationVersion))
	}

	// Verify the first apk
	actualPackageName := apks.Data.Detail.ApplicationVersion[0].Result.Data.Package
	if actualPackageName != packageName {
		t.Errorf("Expected package to be '%s', got %s", packageName, actualPackageName)
	}
}
