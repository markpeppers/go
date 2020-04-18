package xkcdlib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const baseURL = "https://xkcd.com"
const indexDir = "xkcd"

// Download gets the index. Filename will be like 2295.json
func Download(maxComic int) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	indexPath := filepath.Join(homeDir, indexDir)
	files, err := getFileList()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= maxComic; i++ {
		if exists(files, i) {
			continue
		}
		if blacklist(i) {
			continue
		}
		url := fmt.Sprintf("%s/%d/info.0.json", baseURL, i)
		filename := fmt.Sprintf("%d.json", i)
		filepath := filepath.Join(indexPath, filename)
		fmt.Println("Downloading", url)
		err := downloadFile(filepath, url)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func getFileList() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	indexPath := filepath.Join(homeDir, indexDir)
	files, err := ioutil.ReadDir(indexPath)
	if err != nil {
		return nil, err
	}
	filenames := make([]string, 0, len(files))
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func exists(files []string, num int) bool {
	for _, file := range files {
		fileNum, err := strconv.Atoi(strings.Split(file, ".")[0])
		if err != nil {
			log.Fatal(err)
		}
		if fileNum == num {
			return true
		}
	}
	return false
}

func blacklist(i int) bool {
	if i == 404 {
		return true
	}
	return false
}

func downloadFile(filepath, url string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// GetComic gets the string info for the comic comicNum
func GetComic(comicNum int) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filename := filepath.Join(homeDir, indexDir, fmt.Sprintf("%d.json", comicNum))
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
