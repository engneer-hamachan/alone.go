package repository

import "alone/src/domain/model/word"

type WordRepository interface {
	GetFilePaths(root_directory string) ([]string, error)
	GetWords(paths []string) ([]word.Word, error)
}
