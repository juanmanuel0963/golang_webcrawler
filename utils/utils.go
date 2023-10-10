package utils

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

// Function to get the base domain of a given url
func GetBaseDomain(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		fmt.Println("\nInvalid URL")
		os.Exit(1)
	}
	return parts[0] + "//" + parts[2]
}

// Return an URL in the format mydomain_com
func GetShortDomain(inputURL string) string {
	// Parse the input URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("\nInvalid URL:", err)
		return ""
	}

	// Remove "www." from the host if present
	host := strings.TrimPrefix(parsedURL.Hostname(), "www.")

	// Replace periods with underscores to get the base domain
	baseDomain := strings.Replace(host, ".", "_", -1)

	return baseDomain
}
