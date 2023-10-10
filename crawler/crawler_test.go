package crawler

import (
	"reflect"
	"testing"
	"webcrawler/utils"
)

func TestCrawl(t *testing.T) {
	tests := []struct {
		inputURL    string
		expected    map[string]bool
		expectedErr bool
	}{
		{
			inputURL: "https://google.com",
			expected: map[string]bool{
				"https://google.com/services":                                  true,
				"https://google.com/intl/es-419/about.html":                    true,
				"https://google.com/intl/es-419/policies/privacy":              true,
				"https://google.com/intl/es-419/policies/terms":                true,
				"https://google.com/preferences?hl=es-419":                     true,
				"https://google.com/advanced_search?hl=es-419\u0026authuser=0": true,
				"https://google.com/intl/es-419/ads":                           true,
			},
			expectedErr: false,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {

		//Initialize the base domain in the crawler object
		BaseDomain = utils.GetBaseDomain(test.inputURL)

		//Initialize the pages visited map in the crawler object
		CrawledPages.Visited = make(map[string]bool)

		result := Crawl(test.inputURL)

		// Check if the result matches the expected map
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("URL: %s - Expected: %v, Got: %v", test.inputURL, test.expected, result)
		}

		// Check if an error was expected but not received, or vice versa
		if (len(result) == 0 && !test.expectedErr) || (len(result) > 0 && test.expectedErr) {
			t.Errorf("URL: %s - Expected error: %v, Got error: %v", test.inputURL, test.expectedErr, len(result) == 0)
		}
	}
}
