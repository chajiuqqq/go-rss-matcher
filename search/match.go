package search

import (
	"log"
)

// Matcher interface
type Matcher interface {
	Search(feed *Feed, pattern string) ([]*Result, error)
}

// match result struct
type Result struct {
	Field   string
	Content string
}

// match method,根据matcher，pattern，feed进行匹配，并向chan发送结果
func Match(matcher Matcher, pattern string, feed *Feed, resultChan chan *Result) {
	results, err := matcher.Search(feed, pattern)
	if err != nil {
		log.Println(err)
		return
	}
	for _, r := range results {
		resultChan <- r
	}
}

//display，显示结果函数
func Display(resultChan chan *Result) {
	for result := range resultChan {
		log.Printf("%s\n%s", result.Field, result.Content)
	}
}
