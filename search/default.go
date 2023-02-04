package search

// defaultMatcher struct
type defaultMatcher struct{}

// init() 注册matcher
func init() {
	Register("default", &defaultMatcher{})

}

// match()
func (d defaultMatcher) Search(feed *Feed, pattern string) ([]*Result, error) {
	return nil, nil
}
