package seed

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func SeedRoleSAndUsersAndProfles(ctx context.Context, db *sqlx.DB) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	if err = SeedRoles(ctx, tx); err != nil {
		log.Printf("Failed to seed roles: %v", err)
		return err
	}

	if err = SeedUsers(ctx, tx); err != nil {
		log.Printf("Failed to seed users: %v", err)
		return err
	}

	if err = SeedProfiles(ctx, tx); err != nil {
		log.Printf("Failed to seed profiles: %v", err)
		return err
	}

	log.Println("Seeded roles, users, and profiles successfully.")
	return nil
}
