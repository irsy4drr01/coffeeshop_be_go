package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedUsers(ctx context.Context, db *sqlx.Tx) error {
	query := `
		INSERT INTO public.users (id, email, password, role_id, is_verified, is_deleted, created_at, updated_at)
		VALUES
			('6d513ff2-3263-4778-8238-f961acffb44a', 'superadmin@example.com', '$2a$10$g4kXzmt0853QfGgLSW1bu.ImRFHufTTdBtKIhYKpGG/32uCevtUmS', 1, true, false, '2025-06-18 04:47:11.113638+07', NULL),
			('990eb4d2-7658-4624-be54-1fa920f45508', 'admin@example.com', '$2a$10$jj69yP0fg0yl.D78s6LZH.43i3cttmDyxM9VOWSvshOAm5xU3EnBe', 2, true, false, '2025-06-18 11:19:28.886361+07', NULL),
			('deeef659-543e-4dbb-8439-be5df8425cb9', 'user1@example.com', '$2a$10$5r1PLn0HMiqXlR2Ww04.UuyTyPc1/./GBTsCnQQFoeXg.B7IadJlW', 3, true, false, '2025-06-18 11:20:33.964642+07', NULL),
			('2f43d57b-e397-4b0d-9c1a-5f5b482c01ff', 'user2@example.com', '$2a$10$dkn0iV/7TQq4uBklHnEjO.JeZ7ZpuG4STMTK3W3vTOF3vlEQZlj8e', 3, true, false, '2025-06-18 11:20:40.313331+07', NULL),
			('cacec840-1b01-4b31-9774-6cf50ec376f9', 'user3@example.com', '$2a$10$GDw8KmDIWBhD12syl8lxVeDK/7E5yj.Zdeva.lYurOzmbkkQMzCkO', 3, true, false, '2025-06-18 11:20:47.26588+07', NULL),
			('d4dc488a-4aa9-44ac-89e6-a96c4c1480ad', 'user4@example.com', '$2a$10$LNS/dKZcdcBdPfBjOniaXeN1.TeFClByUtD.j3EGoRkp23wMFlAju', 3, true, false, '2025-06-18 13:15:14.323871+07', NULL),
			('88a172fc-1509-400c-9a90-f2e9d7f59b0d', 'user5@example.com', '$2a$10$JP/IUQIrsbfqnws7vTNt6.T7ZEzsXowHPuh3PNt5ECMZhd2.NsGJG', 3, true, false, '2025-06-18 13:15:21.276452+07', NULL)
		ON CONFLICT (email) DO NOTHING;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Failed to seed users: %v", err)
		return err
	}

	log.Println("Seeded users successfully.")
	return nil
}
