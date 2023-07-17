package usecase

import (
	"alone/src/domain/domain_service"
	"alone/src/domain/model/word"
	"alone/src/domain/repository"
)

type SearchAloneUsecase interface {
	GetAloneWords(root_directory string) ([]word.Word, error)
}

type searchAloneUsecase struct {
	wordRepository    repository.WordRepository
	wordDomainService domain_service.WordDomainService
}

func NewSearchAloneUsecase(wr repository.WordRepository, wd domain_service.WordDomainService) SearchAloneUsecase {
	return &searchAloneUsecase{
		wordRepository:    wr,
		wordDomainService: wd,
	}
}

func (su *searchAloneUsecase) GetAloneWords(root_directory string) ([]word.Word, error) {
	file_paths, err := su.wordRepository.GetFilePaths(root_directory)
	if err != nil {
		return nil, err
	}

	words, err := su.wordRepository.GetWords(file_paths)
	if err != nil {
		return nil, err
	}

	words, err = su.wordDomainService.ExtractUniqueWords(words)
	if err != nil {
		return nil, err
	}

	return words, nil
}
