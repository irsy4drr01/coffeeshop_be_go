package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedTax(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO tax (tax_value)
		VALUES
			(0.1)
		ON CONFLICT (id) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed tax: %v", err)
		return err
	}

	log.Println("Seeded tax successfully.")
	return nil
}
