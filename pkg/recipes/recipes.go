package recipes

import (
	"database/sql"
	"log"
)

type Recipe struct {
	ID           int
	Title        string
	Description  string
	Ingredients  string
	Instructions string
}

func ViewRecipe(db *sql.DB, recipe Recipe) {
	query := ``
	_, err := db.Exec(query)
}

func AddRecipe(db *sql.DB, recipe Recipe) error {
	query := `INSERT INTO recipes (title, description, ingredients, instructions) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, recipe.Title, recipe.Description, recipe.Ingredients, recipe.Instructions)
	return err
}

func ListRecipes(db *sql.DB) ([]Recipe, error) {
	rows, err := db.Query("SELECT id, title, description FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []Recipe
	for rows.Next() {
		var r Recipe
		if err := rows.Scan(&r.ID, &r.Title, &r.Description); err != nil {
			log.Println(err)
			continue
		}
		recipes = append(recipes, r)
	}
	return recipes, nil
}
