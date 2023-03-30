package utils

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name     string
        input    string
        expected bool
    }{
        {"Valid URL", "https://www.example.com", true},
        {"Invalid URL - Missing Protocol", "www.example.com", false},
        {"Invalid URL - Missing TLD", "https://www", false},
        {"Invalid URL - Missing Domain", "https://.com", false},
        {"Invalid URL - Invalid Characters", "https://www.!example.com", false},
    }

    for _, tc := range tests {
        tc := tc // capture range variable
        t.Run(tc.name, func(t *testing.T) {
            t.Parallel()
            result := isValidURL(tc.input)
            if result != tc.expected {
                t.Errorf("Expected %v but got %v for input %q", tc.expected, result, tc.input)
            }
        })
    }
}

func TestHashURL(t *testing.T) {
    testCases := []struct {
        name     string
        url      string
        expected string
    }{
        {
            name:     "empty URL",
            url:      "",
            expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", // sha256 of empty string
        },
        {
            name:     "short URL",
            url:      "https://example.org",
            expected: "50d7a905e3046b88638362cc34a31a1ae534766ca55e3aa397951efe653b062b",
        },
        {
            name:     "long URL",
            url:      "https://www.google.com/search?q=go+programming+language+tests&rlz=1C1GCEU_enUS832US832&oq=go+programming+language+tests&aqs=chrome..69i57j0i22i30i457.6409j0j4&sourceid=chrome&ie=UTF-8",
            expected: "3c83abe3571db8877858991ce690bb8db959bf0a57190d2a564dbde14d600c4b",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            hashed := createHash(tc.url)
            if hashed != tc.expected {
                t.Errorf("hashURL(%s) expected %s but got %s", tc.url, tc.expected, hashed)
            }
        })
    }
}

