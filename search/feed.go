package search

import (
	"encoding/json"
	"os"
)

// struct feed
type Feed struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

const dataFile = "data/data.json"

//读取所有feeds，并返回[]*Feed
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	feeds := make([]*Feed, 0)
	err = json.NewDecoder(file).Decode(&feeds)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
