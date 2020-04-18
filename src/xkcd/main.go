package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"xkcdlib"
)

// Rough draft:
// If first arg is "show", 2nd arg should be a number. Show info for that comic number

func main() {
	if os.Args[1] == "show" {
		comicNum, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		err = showComic(comicNum)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}

	matching := xkcdlib.SearchTranscript(os.Args[1])
	fmt.Printf("Matches: %d\n", matching)
}

func showComic(num int) error {
	// Ensure this comic is downloaded
	xkcdlib.Download(num)
	content, err := xkcdlib.GetComic(num)
	if err != nil {
		return err
	}
	fmt.Printf("Comic # %d\n%s\n", num, content)
	return nil
}
