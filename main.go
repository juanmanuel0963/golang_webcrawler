package main

import (
	"fmt"
	"os"
	"sync"
	"time"
	"webcrawler/crawler"
	"webcrawler/trees"
)

func main() {

	//Read the input characters
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run webcrawler.go <starting_url>")
		os.Exit(1)
	}

	//Get the input URL
	startingURL := os.Args[1]

	//Initialize the base domain
	crawler.BaseDomain = crawler.GetBaseDomain(startingURL)

	//Initialize the pages visited
	crawler.CrawledPages.Visited = make(map[string]bool)

	//Create the root node
	root := &trees.TreeNode{Name: startingURL}

	//Create the wait group
	wg := &sync.WaitGroup{}

	//Create the channel
	doneChan := make(chan string)

	//Start the timer
	startTime := time.Now()

	//Increase the wait group count
	wg.Add(1)

	//Call Go routine for crawling recursively an URL
	go crawler.RecursiveCrawl(startingURL, 0, 0, doneChan, root, wg)

	//Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	//Read from doneChan
	for range doneChan {
		// Do nothing here, just let it finish
	}

	//Print the tree
	trees.PrintTree(root, 0)

	//Convert the tree to JSON
	jsonData, err := trees.ConvertTreeToJson(root)
	if err != nil {
		fmt.Printf("Error converting tree to JSON: %v\n", err)
		return
	}

	//Print the JSON data after all crawling is complete
	fmt.Printf("Tree structure in JSON:\n%s\n", jsonData)

	//Finish the timer
	elapsed := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsed)
}
