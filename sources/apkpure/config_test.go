package apkpure_test

import (
	"testing"

	"github.com/pokeguys/apk-crawler/sources/apkpure"
)

func TestApkPureConfig_ValidateConfigValid(t *testing.T) {
	// Test with a valid config
	config := &apkpure.Config{
		SDKVersion: "29",
		Abis:       "arm64-v8a,armeabi-v7a,armeabi",
	}
	message := config.Validate()
	if message != "" {
		t.Errorf("Expected no error, got %s", message)
	}
}

func TestApkPureConfig_ValidateConfigInvalid(t *testing.T) {
	// Test with an invalid config
	config := &apkpure.Config{}
	message := config.Validate()
	if message == "" {
		t.Errorf("Expected an error, got nil")
	}
}

func TestApkPureConfig_Name(t *testing.T) {
	// Test with a valid config
	config := &apkpure.Config{}
	name := config.Name()
	if name == "" {
		t.Errorf("Expected a name, got empty")
	}
}
