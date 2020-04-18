package xkcdlib

import (
	"log"
	"strconv"
	"strings"
)

// Comic represents a single xkcd comic
type Comic struct {
	Month      string
	Num        int
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	ImgURL     string `json:"img"`
	Title      string
	Day        string
}

func loadComics() (map[int]string, error) {
	fileList, err := getFileList()
	if err != nil {
		return nil, err
	}

	comics := make(map[int]string, len(fileList))

	for _, file := range fileList {
		num, err := strconv.Atoi(strings.Split(file, ".")[0])
		if err != nil {
			return nil, err
		}
		content, err := GetComic(num)
		if err != nil {
			return nil, err
		}
		comics[num] = content
	}

	return comics, nil
}

// SearchTranscript searches and returns ints
func SearchTranscript(pattern string) []int {
	comics, err := loadComics()
	if err != nil {
		log.Fatal(err)
	}
	matching := make([]int, 0)
	for key, comic := range comics {
		if strings.Contains(strings.ToLower(comic), strings.ToLower(pattern)) {
			matching = append(matching, key)
		}
	}
	return matching
}
