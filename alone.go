package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// usecase and handler
func main() {
	directory := "./"
	files, err := getRubyFiles(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	wordCounts := make(map[string]int)

	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fileWords := extractWords(string(content))
		for _, word := range fileWords {
			wordCounts[word]++
		}
	}

	uniqueWords := make([]string, 0)
	for word, count := range wordCounts {
		if count == 1 {
			uniqueWords = append(uniqueWords, word)
		}
	}

	for _, word := range uniqueWords {
		fmt.Println(word)
	}
}

// repository
func getRubyFiles(directory string) ([]string, error) {
	var files []string
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".rb" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// domain service
func extractWords(content string) []string {
	var words []string
	r := regexp.MustCompile(`[[:word:]]+`)
	matches := r.FindAllString(content, -1)
	for _, match := range matches {
		words = append(words, match)
	}
	return words
}
