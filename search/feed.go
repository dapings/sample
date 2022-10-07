package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed contains to process a feed info.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds reads and unmarshal the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			_ = file.Close()
		}
	}(file)
	
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	
	return feeds, err
}
