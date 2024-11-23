package ingredients

import (
	"database/sql"
)

type Ingredient struct {
	ID       int
	Name     string
	Quantity float64
	Unit     string
	Expiry   string
}

func AddIngredient(db *sql.DB, ingredient Ingredient) error {
	query := `INSERT INTO pantry (ingredient_name, quantity, unit, expiry_date) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, ingredient.Name, ingredient.Quantity, ingredient.Unit, ingredient.Expiry)
	return err
}

func ListIngredients(db *sql.DB) ([]Ingredient, error) {
	rows, err := db.Query("SELECT id, ingredient_name, quantity, unit, expiry_date FROM pantry")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ingredients []Ingredient
	for rows.Next() {
		var ing Ingredient
		err := rows.Scan(&ing.ID, &ing.Name, &ing.Quantity, &ing.Unit, &ing.Expiry)
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, ing)
	}
	return ingredients, nil
}
