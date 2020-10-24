package utils

import (
	"fmt"
	"net/url"
)

// MakeURLAbsolute makes a given url absolute and returns an error if it was fed an invalid URL
func MakeURLAbsolute(rawurl string) (string, error) {
	parsedURL, err := url.Parse(rawurl)
	if err != nil {
		return "", fmt.Errorf("couldn't parse url: %s, %v", rawurl, err)
	}
	if !parsedURL.IsAbs() {
		parsedURL.Scheme = "http"
	}
	return parsedURL.String(), nil
}
