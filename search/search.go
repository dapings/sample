package search

import (
	"log"
	"sync"
)

// the global map of registered matchers for searching feed.
var matchers = make(map[string]Matcher)

// Register registers a matcher for using by the program.
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

// Run performs the search logic.
func Run(searchItem string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatalln(err)
	}
	
	var wg sync.WaitGroup
	wg.Add(len(feeds))
	
	// an unbuffered channel to receive matched results to display.
	results := make(chan *Result)
	// Launch a goroutine for each feed to find the results.
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers[defaultMatcherName]
		}
		
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchItem, results)
			wg.Done()
		}(matcher, feed)
	}
	
	// Launch a goroutine to monitor when all the work is done.
	go func() {
		wg.Wait()
		
		// close the channel.
		// the signal to the DisplayConsole func for exiting the program.
		close(results)
	}()
	
	DisplayConsole(results)
}