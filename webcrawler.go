package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
	"webcrawler/crawler"
	"webcrawler/file_processing"
	"webcrawler/logger"
	"webcrawler/trees"
	"webcrawler/utils"
)

func main() {

	//Read the input characters
	if len(os.Args) != 3 {
		fmt.Println("\nUsage: go run webcrawler.go <starting_url> <deep_level_to_retrieve>")
		os.Exit(1)
	}

	//Get the input URL
	startingURL := os.Args[1]

	//Get the max deep level to retrieve
	maxDeepLevel, _ := strconv.ParseInt(os.Args[2], 0, 0)

	//Initialize the base domain in the crawler object
	crawler.BaseDomain = utils.GetBaseDomain(startingURL)

	//Initialize the pages visited map in the crawler object
	crawler.CrawledPages.Visited = make(map[string]bool)

	//Start the timer
	startTime := time.Now()

	//Create the root node for saving the sitemap as a tree
	root := &trees.TreeNode{Name: startingURL}

	//Create channel for concurrent retrieval of URLs
	doneChan := make(chan string)

	//Create the wait group to control finalization of goroutines
	wg := &sync.WaitGroup{}

	//Increase the wait group count for control the number of concurrent goroutines
	wg.Add(1)

	//Call to Go routine for crawling recursively an URL
	go crawler.RecursiveCrawl(startingURL, 0, int(maxDeepLevel), 0, doneChan, root, wg)

	//Wait for all goroutines to finish then close the channel
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	//Iterate doneChan
	for range doneChan {
		// Do nothing here, just let it finish
	}

	//Print on the screen the tree with the sitemap
	fmt.Printf("\nTree structure indented:\n")
	trees.PrintTree(root, 0)

	//Convert the tree to JSON
	jsonData, err := trees.ConvertTreeToJson(root)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", crawler.BaseDomain, "Error converting tree to JSON.", err)
		logger.LogError(errMsg)
		return
	}

	//Print on the screen the JSON data after all crawling is complete
	fmt.Printf("\nTree structure in JSON:\n%s\n", jsonData)

	//Get the domain in the format mydomain.com
	domain := utils.GetShortDomain(startingURL)

	// Save JSON data to a file with a timestamp in the filename.
	_, err = file_processing.SaveJSONToFileWithTimestamp(jsonData, domain)
	if err != nil {
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", crawler.BaseDomain, "Error saving JSON to file.", err)
		logger.LogError(errMsg)
		return
	}

	//Finish the timer
	elapsed := time.Since(startTime)
	fmt.Printf("\nTotal time taken: %s\n", elapsed)

}
