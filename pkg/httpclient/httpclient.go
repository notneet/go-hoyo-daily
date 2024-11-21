package httpclient

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/iunary/fakeuseragent"
)

type HttpClient struct {
	client *resty.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: resty.New().SetDebug(os.Getenv("ENV") == "development"),
	}
}

func (h *HttpClient) ApiClientHoyo(cookie, actId, signGame string) *resty.Request {
	randomUA := fakeuseragent.RandomUserAgent()

	return h.client.
		SetHeaders(map[string]string{
			"User-Agent":      randomUA,
			"Accept":          "application/json, text/plain, */*",
			"Accept-Language": "en-US,en;q=0.5",
			"Content-Type":    "application/json;charset=utf-8",
			"Cookie":          cookie,
			"Sec-Fetch-Dest":  "empty",
			"Sec-Fetch-Mode":  "cors",
			"Sec-Fetch-Site":  "same-site",
			"Priority":        "u=0",
			"x-rpc-signgame":  signGame,
		}).
		R().
		SetBody(`{"act_id":"` + actId + `","lang":"en-us"}`)
}
