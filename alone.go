package main

import (
	"alone/infrastructures/persistance"
	"alone/presentation/handler"
	"alone/usecase"
)

func main() {
	wordRepository := persistance.NewWordPersistance()
	searchAloneUsecase := usecase.NewSearchAloneUsecase(wordRepository)
	searchAloneHandler := handler.NewSearchAloneHandler(searchAloneUsecase)

	searchAloneHandler.SearchAlone("./")
}
