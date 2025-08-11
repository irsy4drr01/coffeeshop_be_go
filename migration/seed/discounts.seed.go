package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedDiscounts(ctx context.Context, db *sqlx.DB) error {
	query := `
		INSERT INTO discounts (id, name, discount, expired, is_actived)
		VALUES
			(1, 'Weekend Treats', 0.15, '2025-07-25 16:59:00+07', true),
			(2, 'May Day Sale', 0.2, '2025-07-31 16:59:00+07', true),
			(3, 'Today Sale', 0.2, '2025-07-20 16:59:00+07', true),
			(4, 'Midweek Madness', 0.1, '2025-07-22 23:59:00+07', true),
			(5, 'End of Month Bonanza', 0.25, '2025-07-31 23:59:00+07', true),
			(6, 'Flash Sale', 0.3, '2025-07-19 23:00:00+07', true),
			(7, 'Coffee Lovers Week', 0.12, '2025-07-24 22:00:00+07', true),
			(8, 'Buy 2 Get 1 Week', 0.1667, '2025-07-26 23:59:00+07', true),
			(9, 'Payday Promo', 0.18, '2025-07-27 23:59:00+07', true),
			(10, 'Customer Appreciation', 0.2, '2025-07-31 23:59:00+07', true)
		ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed discounts: %v", err)
		return err
	}

	log.Println("Seeded discounts successfully.")
	return nil
}
