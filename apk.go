package apkcrawler

type Apk struct {
	Name    string `json:"name"`
	Package string `json:"package"`
	Version string `json:"version"`
	Size    int64  `json:"size"`
	URL     string `json:"url"`
	Hash    string `json:"hash"`
}
