package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	flag.Parse()

	rs, err := os.Open(flag.Arg(0)) // pdf
	if err != nil {
		log.Fatalln(err)
	}
	defer rs.Close()

	bs, err := os.ReadFile(flag.Arg(1)) // bookmark
	if err != nil {
		log.Fatalln(err)
	}
	bs, err = addCustomSettings(bs) // <-
	if err != nil {
		log.Fatalln(err)
	}
	rd := bytes.NewBuffer(bs)

	w, err := os.Create(flag.Arg(2)) // output
	if err != nil {
		log.Fatalln(err)
	}
	defer w.Close()

	err = api.ImportBookmarks(rs, rd, w, true, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

type BookmarkTree struct {
	pdfcpu.BookmarkTree
	Settings Settings `json:"settings"`
}

type Settings struct {
	Offset int `json:"offset"`
}

// Add custom settings
func addCustomSettings(bs []byte) ([]byte, error) {
	var t BookmarkTree
	err := json.Unmarshal(bs, &t)
	if err != nil {
		return nil, err
	}

	var f func(b *pdfcpu.Bookmark)
	f = func(b *pdfcpu.Bookmark) {
		b.PageFrom += t.Settings.Offset
		for i := range b.Kids {
			f(&b.Kids[i])
		}
	}

	for i := range t.Bookmarks {
		f(&t.Bookmarks[i])
	}

	return json.Marshal(t.BookmarkTree)
}
