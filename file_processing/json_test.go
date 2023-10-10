package file_processing

import (
	"os"
	"testing"
	"time"
)

func TestSaveJSONToFileWithTimestamp(t *testing.T) {
	// Define a temporary test directory
	testDir := "test_data"
	defer os.RemoveAll(testDir)

	// Create a temporary JSON data
	jsonData := []byte(`{"name": "example", "children": []}`)

	// Test case 1: Successful save
	filename := "example"
	savedname, err := SaveJSONToFileWithTimestamp(jsonData, filename)

	// Check if there was no error
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	time.Sleep(3 * time.Second)

	// Check if the file was created
	_, err = os.Stat("sitemaps/" + savedname)
	if err != nil {
		t.Errorf("Expected file to be created, but got error: %v", err)
	}
	/*
		// Test case 2: Error opening file
		// Create a read-only directory to simulate a permission error
		err = os.MkdirAll(filepath.Join(testDir, "sitemaps"), 0444)
		if err != nil {
			t.Fatal(err)
		}
		_, err = SaveJSONToFileWithTimestamp(jsonData, filename)

		time.Sleep(3 * time.Second)

		// Check if an error occurred
		if err == nil {
			t.Errorf("Expected an error opening the file, but got no error")
		}
	*/
	/*
		// Test case 3: Error writing to file
		// Create a file with read-only permission to simulate a write error
		err = os.WriteFile(filepath.Join(testDir, "sitemaps", filename+"_2022-01-01_12-00-00.json"), []byte("existing data"), 0444)
		if err != nil {
			t.Fatal(err)
		}

		_, err = SaveJSONToFileWithTimestamp(jsonData, filename)

		time.Sleep(3 * time.Second)

		// Check if an error occurred
		if err == nil {
			t.Errorf("Expected an error writing to the file, but got no error")
		}
	*/
}
