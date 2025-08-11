package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedCategories(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO categories (name)
		VALUES 
			('Coffee'),
			('Non-Coffee'),
			('Food'),
			('Dessert'),
			('Snack'),
			('Topping')
		ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed categories: %v", err)
		return err
	}

	log.Println("Seeded categories successfully.")
	return nil
}
