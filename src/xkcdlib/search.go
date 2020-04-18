package xkcdlib

import (
	"encoding/json"
	"log"
	"sort"
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

func loadComics() (map[int]Comic, error) {
	fileList, err := getFileList()
	if err != nil {
		return nil, err
	}

	comics := make(map[int]Comic, len(fileList))

	for _, file := range fileList {
		num, err := strconv.Atoi(strings.Split(file, ".")[0])
		if err != nil {
			return nil, err
		}
		content, err := GetComic(num)
		if err != nil {
			return nil, err
		}
		reader := strings.NewReader(content)
		var result Comic
		err = json.NewDecoder(reader).Decode(&result)
		if err != nil {
			return nil, err
		}
		comics[num] = result
	}

	return comics, nil
}

// SearchTranscript searches and returns ints
func SearchTranscript(pattern string) []Comic {
	comics, err := loadComics()
	if err != nil {
		log.Fatal(err)
	}
	matching := make([]Comic, 0)
	for _, comic := range comics {
		if strings.Contains(strings.ToLower(comic.Transcript), strings.ToLower(pattern)) {
			matching = append(matching, comic)
		}
	}
	sort.Slice(matching, func(i, j int) bool {
		return matching[i].Num < matching[j].Num
	})
	return matching
}
