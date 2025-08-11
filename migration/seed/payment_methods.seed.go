package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedPaymentMethods(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO payment_methods (name)
		VALUES 
			('BRI'),
			('DANA'),
			('BCA'),
			('GoPay'),
			('OVO'),
			('QRIS')
		ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed payment_methods: %v", err)
		return err
	}

	log.Println("Seeded payment_methods successfully.")
	return nil
}
