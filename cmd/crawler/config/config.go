package config

import (
	"flag"
)

type Config struct {
	Source  string
	Package string
	ShowAll bool
}

const (
	defaultSource  = ""
	defaultPackage = ""
	defaultShowAll = false
)

func NewConfig() *Config {
	var source string
	flag.StringVar(&source, "source", defaultSource, "The source to crawl")
	flag.StringVar(&source, "s", defaultSource, "The source to crawl (alias)")

	var packageName string
	flag.StringVar(&packageName, "package", defaultPackage, "The package to crawl")
	flag.StringVar(&packageName, "p", defaultPackage, "The package to crawl (alias)")

	var showAll bool
	flag.BoolVar(&showAll, "all", defaultShowAll, "Get all the versions of the package")
	flag.BoolVar(&showAll, "a", defaultShowAll, "Get all the versions of the package (alias)")
	// TODO: Maybe support for XAPK?
	flag.Parse()

	// Validate the flags
	if source == "" {
		panic("source flag is required")
	}
	if packageName == "" {
		panic("package flag is required")
	}

	return &Config{
		Source:  source,
		Package: packageName,
		ShowAll: showAll,
	}
}
