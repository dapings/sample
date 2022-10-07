package matchers

import (
	"encoding/xml"

	"github.com/dapings/sample/search"
)

func init() {
	var matcher rssMatcher
	search.Register(rssMatcherName, matcher)
}

var (
	rssMatcherName = "rss"
)

type (
	rssMatcher struct{}

	// rssItem defines the item tag fields in the rss document.
	rssItem struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// rssImage defines the image tag fields in the rss document.
	rssImage struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// rssChannel defines the channel tag fields in the rss document.|
	rssChannel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          rssImage `xml:"image"`
		Item           rssItem  `xml:"item"`
	}

	// rssDocument defines the fields with the rss document.
	rssDocument struct {
		XMLName xml.Name   `xml:"rss"`
		Channel rssChannel `xml:"channel"`
	}
)

// retrieve performs a HTTP GET request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	return nil, nil
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, nil
}
