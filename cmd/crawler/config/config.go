package config

import (
	"flag"

	"github.com/pokeguys/apk-crawler/sources"
	"github.com/pokeguys/apk-crawler/sources/apkpure"
)

type Config struct {
	ApkType string
	Source  sources.SourceConfig
	Package string
	ShowAll bool
}

const (
	defaultPackage = ""
	defaultApkType = "APK"
	defaultShowAll = false
)

func NewConfig(args []string) Config {
	var cfg Config
	source := args[0]
	sourceCfg, err := sources.NewSourceConfig(source)
	if err != nil {
		panic(err.Error())
	}
	cfg.Source = sourceCfg

	fs := flag.NewFlagSet(source, flag.ExitOnError)
	registerFlags(fs, &cfg)
	fs.Parse(args[2:])

	// Validate the flags
	message := validateFlags(cfg)
	if message != "" {
		panic(message)
	}

	return cfg
}

func registerFlags(fs *flag.FlagSet, cfg *Config) {
	// Register the common flags
	fs.BoolVar(&cfg.ShowAll, "all", defaultShowAll, "Get all the versions of the package")
	fs.BoolVar(&cfg.ShowAll, "a", defaultShowAll, "Get all the versions of the package (alias)")
	fs.StringVar(&cfg.ApkType, "type", defaultApkType, "The type of the package")
	fs.StringVar(&cfg.ApkType, "t", defaultApkType, "The type of the package (alias)")

	// Register the source-specific flags
	if sourceCfg, ok := cfg.Source.(apkpure.Config); ok {
		apkpure.ParseFlags(fs, &sourceCfg)
	}
}

func validateFlags(cfg Config) string {
	// Validate the package
	var message string
	if cfg.Package == "" {
		message = "package flag is required"
	}

	// Validate the apk type
	if cfg.ApkType == "" {
		message = "apk type flag is required"
	}
	if cfg.ApkType != "APK" && cfg.ApkType != "XAPK" {
		message = "invalid apk type"
	}

	// Validate the source-specific flags
	if message == "" {
		return cfg.Source.Validate()
	}
	return message
}
