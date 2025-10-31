/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long:  `New category...`,
	Run: func(cmd *cobra.Command, args []string) {
		db := GetDB()
		category := GetCategoryDB(db)

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		category.Create(name, description)
		// cmd.Help()
	},
}

func init() {
	categoryCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "a", "", "Nome da categoria")
	newCmd.Flags().StringP("description", "d", "", "Descrição da categoria")
	newCmd.MarkFlagsRequiredTogether("name", "description") //obrigatórios
}
