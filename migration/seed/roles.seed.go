package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedRoles(ctx context.Context, db *sqlx.Tx) error {
	query := `
		INSERT INTO roles (role)
		VALUES 
			('super admin'),
			('admin'),
			('user')
		ON CONFLICT (role) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed roles: %v", err)
		return err
	}

	log.Println("Seeded roles successfully.")
	return nil
}
