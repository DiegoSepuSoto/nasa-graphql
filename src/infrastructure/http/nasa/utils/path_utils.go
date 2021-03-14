package utils

import (
	"net/url"
	"os"
)

func APIKeyParam() url.Values {
	v := url.Values{}
	v.Add("api_key", os.Getenv("NASA_API_KEY"))

	return v
}
