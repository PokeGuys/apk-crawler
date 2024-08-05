package apkpurehttp

import (
	"fmt"
)

const (
	SearchURL = "https://api.pureapk.com/m/v3/cms/app_version?hl=en-US&package_name=%s"
)

func GetSearchURL(packageName string) string {
	return fmt.Sprintf(SearchURL, packageName)
}
