package apkpure

import (
	"flag"
)

const (
	defaultSdkVersion = "29"
	defaultAbis       = "arm64-v8a,armeabi-v7a,armeabi"
)

type Config struct {
	SDKVersion string
	Abis       string
}

func ConfigName() string {
	return "apkpure"
}

func ParseFlags(fs *flag.FlagSet, cfg *Config) {
	fs.StringVar(&cfg.SDKVersion, "sdk-version", defaultSdkVersion, "The SDK version")
	fs.StringVar(&cfg.SDKVersion, "s", defaultSdkVersion, "The SDK version (alias)")
	fs.StringVar(&cfg.Abis, "abis", defaultAbis, "The ABIs")
	fs.StringVar(&cfg.Abis, "b", defaultAbis, "The ABIs (alias)")
}

func (cfg Config) Name() string {
	return ConfigName()
}

func (cfg Config) Validate() string {
	var message string
	if cfg.SDKVersion == "" {
		message = "sdk-version is required"
	}
	if cfg.Abis == "" {
		message = "abis is required"
	}

	return message
}
