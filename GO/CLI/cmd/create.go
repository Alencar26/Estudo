/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		existe, _ := cmd.Flags().GetBool("existe")
		fmt.Printf("Comando existe: %t", existe)

		fmt.Println("\n---------------------")

		quantidade, _ := cmd.Flags().GetInt64("quantidade")
		fmt.Printf("Quantidade informada é: %d", quantidade)

		fmt.Println("\n---------------------")

		//pegando valores das variáveis.
		fmt.Printf("Sua categoria: %s\n", categoria)
		fmt.Printf("Seu tipo: %s\n", tipo)
	},
}

var (
	categoria string
	tipo      string
)

func init() {
	categoryCmd.AddCommand(createCmd)
	categoryCmd.PersistentFlags().String("name", "", "Flag também aparecem em comandos filhos")
	categoryCmd.PersistentFlags().StringP("name2", "n", "SEM NOME", "Flag tipo StringP para ter um comando menor")
	categoryCmd.PersistentFlags().BoolP("existe", "e", true, "Verifica se comando existe.")
	categoryCmd.PersistentFlags().Int64P("quantidade", "q", 0, "Informe um Int64")

	//passando o que receber para uma variável
	categoryCmd.PersistentFlags().StringVarP(&categoria, "categoria", "c", "", "Nome da Categoria")
	categoryCmd.PersistentFlags().StringVarP(&tipo, "tipo", "t", "", "Informe o tipo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
