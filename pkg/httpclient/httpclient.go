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
			"User-Agent":        randomUA,
			"Accept":            "application/json, text/plain, */*",
			"Accept-Language":   "en-US,en;q=0.5",
			"Accept-Encoding":   "gzip, deflate, br, zstd",
			"Content-Type":      "application/json;charset=utf-8",
			"Cookie":            cookie,
			"x-rpc-device_id":   "c5903de3a8735a83da246b8086025512_0.7756426865224945",
			"x-rpc-app_version": "",
			"x-rpc-platform":    "4",
			"x-rpc-device_name": "",
			"x-rpc-client_type": "5",
			"x-rpc-device_fp":   "c5903de3a8735a83da246b8086025512_0.7756426865224945",
			"x-rpc-signgame":    signGame,
			"Origin":            "https://act.hoyolab.com",
			"Connection":        "keep-alive",
			"Referer":           "https://act.hoyolab.com/",
			"Sec-Fetch-Dest":    "empty",
			"Sec-Fetch-Mode":    "cors",
			"Sec-Fetch-Site":    "same-site",
			"Priority":          "u=0",
			"TE":                "trailers",
		}).
		R().
		SetBody(`{"act_id":"` + actId + `","lang":"en-us"}`)
}
