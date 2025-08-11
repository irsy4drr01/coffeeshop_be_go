package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedDeliveryMethods(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO delivery_methods (name, fee)
		VALUES 
			('Dine In', 0),
			('Door Delivery', 15000),
			('Pick Up', 5000)
		ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed delivery methods: %v", err)
		return err
	}

	log.Println("Seeded delivery methods successfully.")
	return nil
}
