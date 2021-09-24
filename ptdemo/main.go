package main

import (
	pt "github.com/monochromegane/the_platinum_searcher"
	"log"
	"os"
)

func main() {
	f := os.Stdout
	search := pt.PlatinumSearcher{
		Err: os.Stderr,
		Out: f,
	}
	search.Run([]string{"-c", "k*", "/Users/i517131/go/src/typedclientdemo"})
	finfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(finfo.Size())
}
