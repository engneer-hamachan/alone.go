package persistance

import (
	"alone/src/domain/domain_service"
	"alone/src/domain/model/word"
)

type wordDomainService struct{}

func NewWordDomainService() domain_service.WordDomainService {
	return &wordDomainService{}
}

func (wd *wordDomainService) ExtractUniqueWords(words []word.Word) ([]word.Word, error) {
	var unique_words []word.Word
	for _, w := range words {
		if w.Count == 1 {
			uw := word.Word{
				Value: w.Value,
				Count: w.Count,
			}
			unique_words = append(unique_words, uw)
		}
	}
	return unique_words, nil
}
