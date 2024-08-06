package sources_test

import (
	"testing"

	"github.com/pokeguys/apk-crawler/sources"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
)

func TestNewCrawlerStrategyValid(t *testing.T) {
	// Test with a valid source
	strategy, err := sources.NewCrawlerStrategy(apkpure.Config{}, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if strategy == nil {
		t.Errorf("Expected a strategy, got nil")
	}
}

func TestNewCrawlerStrategyInvalid(t *testing.T) {
	// Test with an invalid source
	source := "invalid"
	_, err := sources.NewCrawlerStrategy(source, nil)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestNewSourceConfigValid(t *testing.T) {
	// Test with a valid source
	config, err := sources.NewSourceConfig(apkpure.ConfigName())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if config == nil {
		t.Errorf("Expected a config, got nil")
	}
}

func TestNewSourceConfigInvalid(t *testing.T) {
	// Test with an invalid source
	config, err := sources.NewSourceConfig("invalid")
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if config != nil {
		t.Errorf("Expected a nil config, got %v", config)
	}
}
