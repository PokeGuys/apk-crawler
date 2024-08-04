package sources_test

import (
	"testing"

	"github.com/pokeguys/apk-crawler/sources"
)

func TestNewCrawlerStrategyValid(t *testing.T) {
	// Test with a valid source
	source := sources.GetSource(sources.ApkPure)
	strategy := sources.NewCrawlerStrategy(source, nil)
	if strategy == nil {
		t.Errorf("Expected a strategy, got nil")
	}
}

func TestNewCrawlerStrategyInvalid(t *testing.T) {
	// Test with an invalid source
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic, got nil")
		}
	}()
	source := "invalid"
	sources.NewCrawlerStrategy(source, nil)
}
