package search

// the default matcher.
type defaultMatcher struct{}

var defaultType = "default"

func init() {
	var matcher defaultMatcher
	// registers the default matcher.
	Register(defaultType, matcher)
}

// Search implements the behavior for the default matcher.
func (m defaultMatcher) Search(feed *Feed, item string) ([]*Result, error) {
	return nil, nil
}
