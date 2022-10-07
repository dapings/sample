package search

import (
	"log"
)

// Result contains the result of a search op.
type Result struct {
	Filed   string
	Content string
}

// Matcher defines the behavior required by types that want
// to implement a new search type.
type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

// Match concurrently searches for each individual feed as a goroutine.
func Match(matcher Matcher, feed *Feed, searchItem string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchItem)
	if err != nil {
		log.Println(err)
		return
	}
	
	// write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}

// DisplayConsole writes results to the console window.
// the results are received by the individual goroutines.
func DisplayConsole(results chan *Result) {
	// the channel blocks until a result is written to the channel.
	// once the channel is closed, the for loop terminated.
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Filed, result.Content)
	}
}
