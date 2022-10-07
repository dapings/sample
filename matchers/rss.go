package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

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
		XMLName        xml.Name  `xml:"channel"`
		Title          string    `xml:"title"`
		Description    string    `xml:"description"`
		Link           string    `xml:"link"`
		PubDate        string    `xml:"pubDate"`
		LastBuildDate  string    `xml:"lastBuildDate"`
		TTL            string    `xml:"ttl"`
		Language       string    `xml:"language"`
		ManagingEditor string    `xml:"managingEditor"`
		WebMaster      string    `xml:"webMaster"`
		Image          rssImage  `xml:"image"`
		Item           []rssItem `xml:"item"`
	}

	// rssDocument defines the fields with the rss document.
	rssDocument struct {
		XMLName xml.Name   `xml:"rss"`
		Channel rssChannel `xml:"channel"`
	}
)

// retrieve performs a HTTP GET request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided ")
	}

	// retrieve the rss feed document from the web
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// close the response once return from the func.
	defer func(respBody io.ReadCloser) {
		err := respBody.Close()
		if err != nil {
			_ = respBody.Close()
		}
	}(resp.Body)

	// check the status code for a 200
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP Response Error %d\n", resp.StatusCode))
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	var results []*search.Result
	for _, channelItem := range document.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Filed:   "Title",
				Content: channelItem.Title,
			})
		}

		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result{
				Filed:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}
