package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedProfiles(ctx context.Context, db *sqlx.Tx) error {
	query := `
		INSERT INTO public.profiles (id, fullname, phone, address, created_at, updated_at)
		VALUES
			('6d513ff2-3263-4778-8238-f961acffb44a', 'super admin', '', '', '2025-06-18 04:47:11.113638+07', NULL),
			('990eb4d2-7658-4624-be54-1fa920f45508', 'admin', '', '', '2025-06-18 11:19:28.886361+07', NULL),
			('deeef659-543e-4dbb-8439-be5df8425cb9', 'user1', '', '', '2025-06-18 11:20:33.964642+07', NULL),
			('2f43d57b-e397-4b0d-9c1a-5f5b482c01ff', 'user2', '', '', '2025-06-18 11:20:40.313331+07', NULL),
			('cacec840-1b01-4b31-9774-6cf50ec376f9', 'user3', '', '', '2025-06-18 11:20:47.26588+07', NULL),
			('d4dc488a-4aa9-44ac-89e6-a96c4c1480ad', 'user4', '', '', '2025-06-18 13:15:14.323871+07', NULL),
			('88a172fc-1509-400c-9a90-f2e9d7f59b0d', 'user5', '', '', '2025-06-18 13:15:21.276452+07', NULL)
		ON CONFLICT (id) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed profiles: %v", err)
		return err
	}

	log.Println("Seeded profiles successfully.")
	return nil
}
