package main

import (
	pt "github.com/monochromegane/the_platinum_searcher"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./temp.txt")
	search := pt.PlatinumSearcher{
		Err: os.Stderr,
		Out: f,
	}
	search.Run([]string{"-c", "kube", "/Users/i517131/go/src/typedclientdemo"})
	finfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(finfo.Size())

	strs := []string{"os", "kube"}
	for _, str := range strs {
		f, err := os.Create("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
		search := pt.PlatinumSearcher{
			Err: os.Stderr,
			Out: f,
		}
		search.Run([]string{"-c", str, "/Users/i517131/go/src/typedclientdemo"})
		finfo, err := f.Stat()
		log.Println(finfo.Size())
	}
}
