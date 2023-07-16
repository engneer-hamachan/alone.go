package handler

import "alone/usecase"
import "fmt"

type SearchAloneHandler interface {
	SearchAlone(root_directory string)
}

type searchAloneHandler struct {
	searchAloneUsecase usecase.SearchAloneUsecase
}

func NewSearchAloneHandler(su usecase.SearchAloneUsecase) SearchAloneHandler {
	return &searchAloneHandler{
		searchAloneUsecase: su,
	}
}

func (sh *searchAloneHandler) SearchAlone(root_directory string) {
	words, err := sh.searchAloneUsecase.GetAloneWords(root_directory)
	if err != nil {
		error_msg := fmt.Errorf("ErrorMsg:%v", err)
		fmt.Println(error_msg)
		return
	}

	for _, word := range words {
		fmt.Println(word.Value)
	}
}
