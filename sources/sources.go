package sources

type CrawlerSource int

const (
	ApkPure CrawlerSource = iota
)

// Sources is a map of sources.
var Sources = map[CrawlerSource]string{
	ApkPure: "apkpure",
}

func GetSources() []string {
	sources := make([]string, len(Sources))
	for source, name := range Sources {
		sources[source] = name
	}
	return sources
}

// GetSource returns the source name.
func GetSource(source CrawlerSource) string {
	return Sources[source]
}

// GetSourceByName returns the source by name.
func GetSourceByName(name string) CrawlerSource {
	for source, n := range Sources {
		if n == name {
			return source
		}
	}
	return -1
}
