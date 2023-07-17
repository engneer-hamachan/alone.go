package word

type Word struct {
	Value string
	Count int
}

type Words struct {
	Words []Word
}

func (w Words) ExtractUniqueWords() Words {
	var unique_words Words
	for _, word := range w.Words {
		if word.Count == 1 {
			uw := Word{
				Value: word.Value,
				Count: word.Count,
			}
			unique_words.Words = append(unique_words.Words, uw)
		}
	}
	return unique_words
}
