package utils

import (
	"reflect"
	"testing"
)

func TestMakeURLAbsolute(t *testing.T) {
	t.Run("leaves a url that is already absolute, absolute", func(t *testing.T) {
		testURLs := []string{
			"https://www.google.com",
			"http://canrozanes.com",
			"https://canrozanes.com",
			"http://www.canrozanes.com",
		}
		outputtedURLs := []string{}
		for _, testURL := range testURLs {
			url, _ := MakeURLAbsolute(testURL)
			outputtedURLs = append(outputtedURLs, url)
		}
		if !reflect.DeepEqual(testURLs, outputtedURLs) {
			t.Errorf("got %v, want %v", outputtedURLs, testURLs)
		}
	})
	t.Run("makes a url absolute if it wasn't before", func(t *testing.T) {
		url := "canrozanes.com"

		got, _ := MakeURLAbsolute(url)
		want := "http://canrozanes.com"

		AssertString(t, got, want)
	})
	t.Run("throws error if it isn't a valid url", func(t *testing.T) {
		url := " _:/.com"

		_, err := MakeURLAbsolute(url)
		want := "couldn't parse url:  _:/.com, parse \" _:/.com\": first path segment in URL cannot contain colon"

		AssertError(t, err, want)
	})
}
