package search

import (
	"fmt"
	"log"
	"sync"
)

var matchMap map[string]Matcher

func Run(pattern string) {
	//读取所有feeds
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//创建waitGroup
	var wg sync.WaitGroup
	wg.Add(len(feeds))

	//chan
	results := make(chan *Result)

	//遍历feeds，给每个feed寻找matcher，并开启goroutine进行match
	for _, feed := range feeds {
		matcher, ok := matchMap[feed.Type]
		if !ok {
			matcher = matchMap["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			defer wg.Done()
			Match(matcher, pattern, feed, results)
		}(matcher, feed)
	}

	//goroutine监听match任务是否完成，并关闭chan
	go func() {
		wg.Wait()
		close(results)
	}()

	//展示结果
	Display(results)
}

//注册matcher

func Register(matchType string, matcher Matcher) error {
	if matchMap == nil {
		matchMap = make(map[string]Matcher)
	}
	//已注册
	if _, ok := matchMap[matchType]; ok {
		return fmt.Errorf("%s 无法重复注册", matchType)
	}
	matchMap[matchType] = matcher
	return nil
}
