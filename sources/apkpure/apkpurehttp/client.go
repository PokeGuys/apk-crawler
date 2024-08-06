package apkpurehttp

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"google.golang.org/protobuf/proto"

	pb "github.com/pokeguys/apk-crawler/proto"
)

type Client interface {
	GetVersions(sdkVersion, abis, packageName string) (*pb.ApkPureResponse, error)
}

type client struct {
	client *resty.Client
}

func NewClient(c *resty.Client) Client {
	return &client{
		client: c,
	}
}

func (c *client) GetVersions(sdkVersion, abis, packageName string) (*pb.ApkPureResponse, error) {
	resp, err := c.client.R().
		SetHeaders(map[string]string{
			"x-cv":   "3172501",
			"x-sv":   sdkVersion,
			"x-abis": abis,
			"x-gp":   "1",
		}).
		Get(GetSearchURL(packageName))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.RawBody().Close()

	var apiResult pb.ApkPureResponse
	err = proto.Unmarshal(resp.Body(), &apiResult)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &apiResult, nil
}
