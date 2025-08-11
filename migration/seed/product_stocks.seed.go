package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedProductStocks1(ctx context.Context, db *sqlx.DB) error {
	query := `
		WITH size_options AS (
			SELECT unnest(array[1, 2, 3]) AS size_id
		)
		INSERT INTO product_stocks (product_id, size_id, stock)
		SELECT 
			p.id,
			s.size_id,
			100 AS stock
		FROM products p
		CROSS JOIN size_options s
		WHERE p.category_id IN (1, 2)
		ON CONFLICT (product_id, size_id) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed product_stocks 1: %v", err)
		return err
	}

	log.Println("Seeded product_stocks 1 successfully.")
	return nil
}

func SeedProductStocks2(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO product_stocks (product_id, size_id, stock)
		SELECT p.id, 4, 100 AS stock
		FROM products p
		WHERE p.category_id IN (3, 4, 5, 6)
		ON CONFLICT (product_id, size_id) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed product_stocks 2: %v", err)
		return err
	}

	log.Println("Seeded product_stocks 2 successfully.")
	return nil
}

func SeedProductStocks(ctx context.Context, db *sqlx.DB) error {
	if err := SeedProductStocks1(ctx, db); err != nil {
		return err
	}
	if err := SeedProductStocks2(ctx, db); err != nil {
		return err
	}
	return nil
}
