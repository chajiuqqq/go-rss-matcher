package matcher

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/chajiuqqq/rss-matcher/search"
)

func init() {
	search.Register("rss", &rssMatcher{})
}

// struct:rssDocument,channel,image,item
type (
	rssDocument struct {
		Channel channel `xml:"channel"`
	}
	channel struct {
		Title       string  `xml:"title"`
		Link        string  `xml:"link"`
		Image       image   `xml:"image"`
		Description string  `xml:"description"`
		Items       []*item `xml:"item"`
	}
	item struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
	}
	image struct {
		Title string `xml:"title"`
		Url   string `xml:"url"`
		Link  string `xml:"link"`
	}
)

// rssMatcher struct
type rssMatcher struct {
}

// search，在feeds中遍历搜索pattern，并返回[]*Result

func (d rssMatcher) Search(feed *search.Feed, pattern string) ([]*search.Result, error) {
	var results []*search.Result
	rssDocument, err := retrieve(feed)
	if err != nil {
		return nil, err
	}
	for _, item := range rssDocument.Channel.Items {
		//匹配title
		matched, err := regexp.MatchString(pattern, item.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{Field: "title", Content: item.Title})
		}
		//匹配description
		matched, err = regexp.MatchString(pattern, item.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{Field: "description", Content: item.Description})
		}
	}
	return results, nil
}

// retrieve 网络读取网址的rss xml文件， 并转为rssDocument
func retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.Link == "" {
		return nil, errors.New("no rss feed uri provided")
	}
	resp, err := http.Get(feed.Link)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error: %s code %d", feed.Link, resp.StatusCode)
	}
	rssDocument := new(rssDocument)
	err = xml.NewDecoder(resp.Body).Decode(rssDocument)
	return rssDocument, err
}
