package file_processing

import (
	"fmt"
	"os"
	"time"
	"webcrawler/crawler"
	"webcrawler/logger"
)

// Saves JSON data to a file with a timestamp in the filename.
func SaveJSONToFileWithTimestamp(jsonData []byte, prefix string) (string, error) {
	// Create the "sitemaps" subfolder if it doesn't exist
	sitemapsFolder := "sitemaps"
	if _, err := os.Stat(sitemapsFolder); os.IsNotExist(err) {
		os.Mkdir(sitemapsFolder, os.ModePerm)
	}

	// Set the format expected for output
	timestamp := time.Now().Format("2006-01-02_15-04-05") // Format: YYYY-MM-DD_HH-MM-SS

	// Create the filename within the "sitemaps" subfolder
	filename := fmt.Sprintf("%s/%s_%s.json", sitemapsFolder, prefix, timestamp)

	// Open the file on disk
	file, err := os.Create(filename)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", crawler.BaseDomain, "Error opening the file", filename)
		logger.LogError(errMsg)
		return filename, err
	}
	defer file.Close()

	// Write data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", crawler.BaseDomain, "Error writing to the file", filename)
		logger.LogError(errMsg)
		return filename, err
	}
	fmt.Printf("\nJSON data saved to %s\n", filename)

	return filename, nil
}
