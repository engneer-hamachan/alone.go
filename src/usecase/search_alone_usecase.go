package usecase

import (
	"alone/src/domain/model/word"
	"alone/src/domain/repository"
)

type SearchAloneUsecase interface {
	GetAloneWords(root_directory string) (*word.Words, error)
}

type searchAloneUsecase struct {
	wordRepository repository.WordRepository
}

func NewSearchAloneUsecase(wr repository.WordRepository) SearchAloneUsecase {
	return &searchAloneUsecase{
		wordRepository: wr,
	}
}

func (su *searchAloneUsecase) GetAloneWords(root_directory string) (*word.Words, error) {
	file_paths, err := su.wordRepository.GetFilePaths(root_directory)
	if err != nil {
		return nil, err
	}

	words, err := su.wordRepository.GetWords(file_paths)
	if err != nil {
		return nil, err
	}

	unique_words := words.ExtractUniqueWords()

	return &unique_words, nil
}
