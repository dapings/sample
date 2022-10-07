package search

// the default matcher.
type defaultMatcher struct{}

var defaultMatcherName = "default"

func init() {
	var matcher defaultMatcher
	// registers the default matcher.
	Register(defaultMatcherName, matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, item string) ([]*Result, error) {
	return nil, nil
}
