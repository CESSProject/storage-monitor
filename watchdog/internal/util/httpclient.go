package util

import (
	"fmt"
	"github.com/CESSProject/watchdog/constant"
	"github.com/go-resty/resty/v2"
	"time"
)

type HTTPClient struct {
	RestyDefaultHttpClient *resty.Client
}

func NewHTTPClient() *HTTPClient {
	client := resty.New().
		SetRetryCount(constant.HttpMaxRetry).
		SetRetryWaitTime(constant.HttpRetryWaitTime * time.Second).
		SetTimeout(constant.HttpTimeout * time.Second)
	return &HTTPClient{RestyDefaultHttpClient: client}
}

func (c *HTTPClient) Request(method, url string, body interface{}, result interface{}) error {
	req := c.RestyDefaultHttpClient.R().
		SetBody(body).
		SetResult(result)

	resp, err := req.Execute(method, url)
	if err != nil {
		return fmt.Errorf("requset failed: %v", err)
	}

	if resp.IsError() {
		return fmt.Errorf("API returned error code: %d", resp.StatusCode())
	}

	return nil
}

func (c *HTTPClient) Get(url string, result interface{}) error {
	return c.Request("GET", url, nil, result)
}

func (c *HTTPClient) Post(url string, body interface{}, result interface{}) error {
	return c.Request("POST", url, body, result)
}
