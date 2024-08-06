package apkpurehttp

import (
	pb "github.com/pokeguys/apk-crawler/proto"
)

func MockApkPureResponse(packageName, apkType string, count int) *pb.ApkPureResponse {
	apiResult := &pb.ApkPureResponse{
		Data: &pb.ApkPureResponseData{
			Detail: &pb.ApkPureResponseDetail{},
		},
	}
	for i := 0; i < count; i++ {
		apiResult.Data.Detail.ApplicationVersion = append(
			apiResult.Data.Detail.ApplicationVersion,
			&pb.ApkPureApplicationVersion{
				Result: &pb.ApkPureApplicationSearchResult{
					Data: &pb.ApkPureApplicationVersionData{
						Name:         "Example App",
						DisplayName:  "Example App",
						Package:      packageName,
						MinorVersion: "1.0",
						Version:      "1.0.0",
						Hash:         "1234567890",
						Description:  "This is an example app",
						PatchNotes:   "This is an example app",
						Status:       "Published",
						Developer:    "Example Developer",
						Download: &pb.ApkPureApplicationDownload{
							Name:       "Example App",
							Sha1:       "1234567890",
							Size:       1024,
							TorrentUrl: "https://example.com/example.apk",
							TrackerUrl: "https://example.com/example.apk",
							Type:       apkType,
							Url:        "https://example.com/example.apk",
						},
					},
				},
			},
		)
	}
	return apiResult
}
