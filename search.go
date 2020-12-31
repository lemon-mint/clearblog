package main

import (
	"os"
	"sync"

	"github.com/blevesearch/bleve"
)

func isexists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

var searchLock sync.Mutex

var indexer bleve.Index = func() bleve.Index {
	searchLock.Lock()
	defer searchLock.Unlock()
	if isexists("post.bleve") {
		index, err := bleve.Open("post.bleve")
		if err != nil {
			os.Exit(-1)
		}
		return index
	}
	index, err := bleve.New("post.bleve", bleve.NewIndexMapping())
	if err != nil {
		os.Exit(-1)
	}
	return index
}()
