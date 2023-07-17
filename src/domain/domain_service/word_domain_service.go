package domain_service

import "alone/src/domain/model/word"

type WordDomainService interface {
	ExtractUniqueWords(words []word.Word) ([]word.Word, error)
}
