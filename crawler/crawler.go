package crawler

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"webcrawler/logger"
	"webcrawler/trees"

	"golang.org/x/net/html"
)

// Struct to control the visited pages
type Pages struct {
	VisitedMutex sync.Mutex
	Visited      map[string]bool
}

// Var to save the base domain
var BaseDomain = ""

// Var to control the visited pages
var CrawledPages Pages

// Go routine for crawling recursively an URL.
// The sitemap is stored in the "parentNode" object as a tree with head nodes and sons
func RecursiveCrawl(parentURL string, currentLevel int, maxLevel int, parent int, doneChan chan string, parentNode *trees.TreeNode, wg *sync.WaitGroup) {
	//Defers the closing of the parent channel
	defer func() {
		doneChan <- parentURL
	}()

	if currentLevel < maxLevel {

		//Retrieve the list of links in the URL and save them in a map
		foundLinks := Crawl(parentURL)

		//Print the current level and parent URL
		fmt.Printf("\nLevel:%v, Parent: %v, %s", currentLevel, parent, parentURL)
		//Counter of links found
		counter := 0

		//Iterate the map of links found
		for sonURL := range foundLinks {

			//Increase the number of links found
			counter++

			//fmt.Printf("Son %v: %s\n", counter, sonURL)

			// Create son node
			sonNode := &trees.TreeNode{Name: sonURL}

			// Add child node to parent
			parentNode.Children = append(parentNode.Children, sonNode)

			// Increment the wait group count as we are invoquing a new goroutine
			wg.Add(1)

			//Call Go routine for crawling recursively the son URL
			go RecursiveCrawl(sonURL, currentLevel+1, maxLevel, counter, doneChan, sonNode, wg)
		}
	}

	// Decrement the wait group count
	wg.Done()
}

// Craw an URL to get the links
func Crawl(parentURL string) map[string]bool {

	//Mutex object for controling read/write access to CrawledPages object
	CrawledPages.VisitedMutex.Lock()
	defer CrawledPages.VisitedMutex.Unlock()

	//Create map to sabe found links in the page
	foundLinks := make(map[string]bool, 0)

	//If the page was visited previously, stop processing
	if CrawledPages.Visited[parentURL] {
		return foundLinks
	}

	//Mark page as visited
	CrawledPages.Visited[parentURL] = true

	//Get the content of the remote URL
	response, err := http.Get(parentURL)

	//Error fetching the URL
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", parentURL, err)
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", BaseDomain, "Error fetching "+parentURL, err)
		logger.LogError(errMsg)

		return foundLinks
	}

	//Set the closing of the body at end of the function
	defer response.Body.Close()

	//URL invocation not successfull
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error fetching %s: Status Code %d\n", parentURL, response.StatusCode)
		errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", BaseDomain, "Error fetching "+parentURL+" Status Code "+fmt.Sprint(response.StatusCode), err)
		logger.LogError(errMsg)

		return foundLinks
	}

	//Tokenie the response body
	tokenizer := html.NewTokenizer(response.Body)

	//Iterate the tokens
	for {

		//Fetch next token
		tokenType := tokenizer.Next()

		//Evaluate the token type
		switch tokenType {

		//Error token or End Of File
		case html.ErrorToken:

			//Get the error
			err := tokenizer.Err()

			//If is End of File
			if err == io.EOF {
				//Finish processing
				return foundLinks
			}

			errMsg := fmt.Sprintf("Domain: %s Message: %s Error: %s", BaseDomain, "Error tokenizing HTML ", tokenizer.Err())
			logger.LogError(errMsg)
		//Start tag token
		case html.StartTagToken, html.SelfClosingTagToken:

			//Get the token
			token := tokenizer.Token()

			//Token type "a"
			if token.Data == "a" {

				//Iterate the attributes
				for _, attr := range token.Attr {

					//If href attribute found
					if attr.Key == "href" {

						//Get the found link
						foundLink := attr.Val

						//If len of link is > 0
						if len(foundLink) > 0 {

							//Trims the "/" character at end of the link
							if string(foundLink[len(foundLink)-1]) == "/" {
								foundLink = foundLink[:len(foundLink)-1]
							}

							//If link is relative
							if strings.HasPrefix(foundLink, "/") {
								foundLink = parentURL + foundLink
							}

							//If not sublink of link to another domain and different to parent url
							if strings.HasPrefix(foundLink, BaseDomain) && foundLink != parentURL {
								foundLinks[foundLink] = true
							}
						}
					}
				}
			}
		}
	}
}
