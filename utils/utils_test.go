package utils

import (
	"testing"
)

func TestGetBaseDomain(t *testing.T) {
	// Test case 1: Valid URL with http
	url1 := "http://example.com"
	baseDomain1 := GetBaseDomain(url1)
	expectedDomain1 := "http://example.com"
	if baseDomain1 != expectedDomain1 {
		t.Errorf("Expected base domain: %s, but got: %s", expectedDomain1, baseDomain1)
	}

	// Test case 2: Valid URL with https
	url2 := "https://www.example.com/path/to/resource"
	baseDomain2 := GetBaseDomain(url2)
	expectedDomain2 := "https://www.example.com"
	if baseDomain2 != expectedDomain2 {
		t.Errorf("Expected base domain: %s, but got: %s", expectedDomain2, baseDomain2)
	}

	// Test case 3: URL with only the domain
	url3 := "example.com"
	baseDomain3 := GetBaseDomain(url3)
	expectedDomain3 := "example.com"
	if baseDomain3 != expectedDomain3 {
		t.Errorf("Expected base domain: %s, but got: %s", expectedDomain3, baseDomain3)
	}

	// Test case 4: Invalid URL (no slashes)
	url4 := "invalid-url"
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected GetBaseDomain to panic for an invalid URL, but it didn't")
		}
	}()
	GetBaseDomain(url4)
}

func TestGetShortDomain(t *testing.T) {
	// Test case 1: Valid URL with http
	url1 := "http://example.com"
	shortDomain1 := GetShortDomain(url1)
	expectedShortDomain1 := "example_com"
	if shortDomain1 != expectedShortDomain1 {
		t.Errorf("Expected short domain: %s, but got: %s", expectedShortDomain1, shortDomain1)
	}

	// Test case 2: Valid URL with https
	url2 := "https://www.example.com/path/to/resource"
	shortDomain2 := GetShortDomain(url2)
	expectedShortDomain2 := "www_example_com"
	if shortDomain2 != expectedShortDomain2 {
		t.Errorf("Expected short domain: %s, but got: %s", expectedShortDomain2, shortDomain2)
	}

	// Test case 3: URL with only the domain
	url3 := "example.com"
	shortDomain3 := GetShortDomain(url3)
	expectedShortDomain3 := "example_com"
	if shortDomain3 != expectedShortDomain3 {
		t.Errorf("Expected short domain: %s, but got: %s", expectedShortDomain3, shortDomain3)
	}

	// Test case 4: Invalid URL (no slashes)
	url4 := "invalid-url"
	shortDomain4 := GetShortDomain(url4)
	if shortDomain4 != "" {
		t.Errorf("Expected an empty short domain for an invalid URL, but got: %s", shortDomain4)
	}
}
