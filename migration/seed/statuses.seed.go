package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedStatuses(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO statuses (status)
		VALUES 
			('Pending'),
			('Processing'),
			('Completed'),
			('Cancelled')
		ON CONFLICT (status) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed statuses: %v", err)
		return err
	}

	log.Println("Seeded statuses successfully.")
	return nil
}
