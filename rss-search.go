package main

import (
	"log"
	"os"

	"github.com/chajiuqqq/rss-matcher/search"

	_ "github.com/chajiuqqq/rss-matcher/matcher"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("China")
}
