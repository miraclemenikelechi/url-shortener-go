package utils

import "net/url"

func IsValidURL(text string) bool {
	parsedURL, err := url.Parse(text)
	if err != nil {
		return false
	}

	if parsedURL.Host == "" {
		return false
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}

	return true
}
