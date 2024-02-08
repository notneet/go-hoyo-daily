package helper

import (
	"github.com/go-resty/resty/v2"
	"github.com/iunary/fakeuseragent"
)

func ApiClientHoyo(host string, cookie string, actId string) *resty.Request {
	randomUA := fakeuseragent.RandomUserAgent()

	return resty.New().
		SetHeaders(map[string]string{
			"User-Agent":      randomUA,
			"Accept":          "application/json, text/plain, */*",
			"Accept-Language": "en-US,en;q=0.5",
			"Content-Type":    "application/json;charset=utf-8",
			"Cookie":          cookie,
			"Sec-Fetch-Dest":  "empty",
			"Sec-Fetch-Mode":  "cors",
			"Sec-Fetch-Site":  "same-site",
		}).
		SetBaseURL(host).
		R().
		SetBody(`{"act_id":"` + actId + `","lang":"en-us"}`)
}
