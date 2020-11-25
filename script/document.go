package script

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

type Document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func LoadDocuments(path string) ([]Document, error) {
	// open file
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// decompress
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	// read xml
	dec := xml.NewDecoder(gz)
	dump := struct {
		Docs []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	// dump
	docs := dump.Docs
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
