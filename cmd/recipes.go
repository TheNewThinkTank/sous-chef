/*
Copyright © 2024 Gustav Rasmussen
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"log"
	"sous-chef/internal/db"
	"sous-chef/pkg/recipes"
)

// recipesCmd represents the recipes command
var recipesCmd = &cobra.Command{
	Use:   "recipes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recipes called")
	},
}

func init() {
	rootCmd.AddCommand(recipesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recipesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recipesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func viewRecipe(id int) {
	db.InitDB("sous_chef.db")
	defer db.DB.Close()

	recipe, err := recipes.GetRecipe(db.DB, id)
	if err != nil {
		log.Fatalf("Error retrieving recipe: %v", err)
	}

	fmt.Printf("Title: %s\nDescription: %s\nIngredients: %s\nInstructions: %s\n",
		recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions)
}

func listRecipes() {
	db.InitDB("sous_chef.db")
	defer db.DB.Close()

	recipeList, err := recipes.ListRecipes(db.DB)
	if err != nil {
		log.Fatalf("Error retrieving recipes: %v", err)
	}

	fmt.Println("Recipes:")
	for _, recipe := range recipeList {
		fmt.Printf("ID: %d, Title: %s, Description: %s\n", recipe.ID, recipe.Title, recipe.Description)
	}
}
