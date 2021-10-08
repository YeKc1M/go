package main

import (
	pt "github.com/monochromegane/the_platinum_searcher"
	"log"
	"os"
)

func main() {
	f, err := os.Create("./temp.txt")
	defer func() {
		err := os.Remove("./temp.txt")
		if err != nil {
			log.Fatal(err)
		}
	}()
	search := pt.PlatinumSearcher{
		Err: os.Stderr,
		Out: f,
	}
	search.Run([]string{"-c", "-e", "(kube)|(os)", "/Users/i517131/go/src/typedclientdemo"})
	finfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(finfo.Size())

	strs := []string{"os", "kube", "231456"}
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
