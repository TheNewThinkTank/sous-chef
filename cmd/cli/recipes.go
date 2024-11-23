package main

import (
	"fmt"
	"log"
	"sous-chef/internal/db"
	"sous-chef/pkg/recipes"
)

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
