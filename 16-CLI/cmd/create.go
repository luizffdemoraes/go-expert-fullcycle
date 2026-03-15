/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
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
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetDB()
		defer db.Close()
		categoryDB := GetCategoryDB(db)
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")

		cat, err := categoryDB.Create(name, description)
		if err != nil {
			return fmt.Errorf("criar categoria: %w", err)
		}
		fmt.Printf("Categoria criada: id=%s name=%s description=%s\n", cat.ID, cat.Name, cat.Description)
		return nil
	},
}

func init() {
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "", "Name of the category")
	createCmd.Flags().StringP("description", "d", "", "Description of the category")
	createCmd.MarkFlagRequired("description")
}
