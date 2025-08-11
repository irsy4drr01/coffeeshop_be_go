package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedSizes(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO sizes (name, additional_price)
		VALUES 
			('Reguler', 0),
			('Medium', 0.25),
			('Large', 0.5),
			('-', 0)
		ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed sizes: %v", err)
		return err
	}

	log.Println("Seeded sizes successfully.")
	return nil
}
