package main

import (
	"fullTextSearch/script"
	"log"
	"time"
)

const (
	dataPath string = "./data/enwiki-latest-abstract1.xml.gz"
	queryStr string = "Anarchism"
)

func main() {
	// load docs
	start := time.Now()
	docs, err := script.LoadDocuments(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	// create index
	start = time.Now()
	index := make(script.Index)
	index.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	// query
	start = time.Now()
	matchedIDs := index.Search(queryStr)
	log.Printf("Search found %d documents in %v", len(matchedIDs), time.Since(start))

	for _, id := range matchedIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
