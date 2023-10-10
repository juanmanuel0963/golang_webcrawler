package file_processing

import (
	"fmt"
	"os"
	"time"
)

// Save JSON data to a file with a timestamp in the filename.
func SaveJSONToFileWithTimestamp(jsonData []byte, prefix string) error {

	//Set the format expected for output
	timestamp := time.Now().Format("2006-01-02_15-04-05") // Format: YYYY-MM-DD_HH-MM-SS

	//Create the filename
	filename := fmt.Sprintf("%s_%s.json", prefix, timestamp)

	//Open file on disc
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write data to file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	fmt.Printf("\nJSON data saved to %s\n", filename)
	return nil
}
