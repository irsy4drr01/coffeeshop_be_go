package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedDiscountProducts(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO discount_products (product_id, discount_id)
		VALUES  
			('4e841656-596c-434d-b5bb-1f27f5d7418c', 1),
			('5740f5fe-5178-4933-9e7b-63cab72fd79a', 1),
			('e6622f29-2a8b-4110-9a82-8616eed29570', 1),
			('7af264d9-fa31-45a4-8948-b2db4c267fd6', 2),
			('f866f4f6-f89c-4395-90ca-241dfb52951c', 2),
			('287ec09d-928c-4562-9f29-86ad95dce6f6', 2)
		ON CONFLICT (product_id, discount_id) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed discount_products: %v", err)
		return err
	}

	log.Println("Seeded discount_products successfully.")
	return nil
}
