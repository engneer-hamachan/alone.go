/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"alone/infrastructures/persistance"
	"alone/presentation/handler"
	"alone/usecase"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alone",
	Short: "search alone words",
	Long: `
    search alone words for your project
  `,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := cmd.Flags().GetString("dir")
		if err != nil {
			fmt.Println(err)
			return
		}

		wordRepository := persistance.NewWordPersistance()
		searchAloneUsecase := usecase.NewSearchAloneUsecase(wordRepository)
		searchAloneHandler := handler.NewSearchAloneHandler(searchAloneUsecase)
		searchAloneHandler.SearchAlone(dir)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("dir", "d", "", "Search Directory Name")
}
