package persistance

import (
	"alone/src/domain/model/word"
	"alone/src/domain/repository"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type wordPersistance struct{}

func NewWordPersistance() repository.WordRepository {
	return &wordPersistance{}
}

func (wp *wordPersistance) GetFilePaths(root_directory string) ([]string, error) {
	var files []string

	err := filepath.Walk(root_directory, func(path string, info os.FileInfo, err error) error {
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

func (wp *wordPersistance) GetWords(paths []string) ([]word.Word, error) {
	var words []word.Word
	counts := make(map[string]int)

	for _, path := range paths {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			continue
		}

		var results []string

		r := regexp.MustCompile(`[[:word:]]+`)

		matches := r.FindAllString(string(content), -1)
		for _, match := range matches {
			results = append(results, match)
		}

		for _, value := range results {
			counts[value]++
		}
	}

	for value, count := range counts {
		if count == 1 {
			word := word.Word{
				Value: value,
				Count: count,
			}
			words = append(words, word)
		}
	}
	return words, nil
}
