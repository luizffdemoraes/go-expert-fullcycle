/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		fmt.Println("Category called with name:", name)
		exists, _ := cmd.Flags().GetBool("exists")
		fmt.Println("Category exists:", exists)
		id, _ := cmd.Flags().GetInt("id")
		fmt.Println("Category ID:", id)
	},
	// hook para executar antes da execução do comando
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado antes da execução do Run")
	},
	// hook para executar depois da execução do comando
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Chamado depois da execução do Run")
	},
	// hook para executar em caso de erro
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("Ocorreu um erro")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "Y", "Name of the category")
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if the category exists")
	categoryCmd.PersistentFlags().IntP("id", "i", 0, "ID of the category")
	// categoryCmd.PersistentFlags().String("name", "", "Name of the category") // flag global
	// categoryCmd.Flags().String("name", "", "Name of the category") // flag local

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
