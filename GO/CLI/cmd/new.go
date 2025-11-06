/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Alencar26/Estudo/GO/CLI/internal/database"
	"github.com/spf13/cobra"
)

func newCreadCmd(categoryDB database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "New caretorry",
		Long:  "New category",
		RunE:  runCreate(categoryDB),
	}
}

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {

		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		_, err := categoryDB.Create(name, description)
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	newCmd := newCreadCmd(GetCategoryDB(GetDB()))
	categoryCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("name", "a", "", "Nome da categoria")
	newCmd.Flags().StringP("description", "d", "", "Descrição da categoria")
	newCmd.MarkFlagsRequiredTogether("name", "description") //obrigatórios
}
