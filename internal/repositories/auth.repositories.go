package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepoInterface interface {
	CreateUserAndProfile(ctx context.Context, user *models.UserAuth, profile *models.ProfileAuth) (*models.UserAuth, *models.ProfileAuth, error)
	GetByEmail(ctx context.Context, email string) (*models.UserAuth, error)
}

type RepoAuth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *RepoAuth {
	return &RepoAuth{db: db}
}

func (r *RepoAuth) CreateUserAndProfile(ctx context.Context, user *models.UserAuth, profile *models.ProfileAuth) (*models.UserAuth, *models.ProfileAuth, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// Email validation
	var exists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM public.users WHERE email = $1 AND is_deleted = false);`
	if err = tx.QueryRowxContext(ctx, checkQuery, user.Email).Scan(&exists); err != nil {
		return nil, nil, err
	}
	if exists {
		return nil, nil, fmt.Errorf("email already in use")
	}

	// Insert user
	userQuery := `
		INSERT INTO public.users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, password, created_at;
	`
	row := tx.QueryRowxContext(ctx, userQuery, user.Email, user.Password)
	if err = row.StructScan(user); err != nil {
		return nil, nil, err
	}

	// Insert profile
	profile.ID = user.ID
	profileQuery := `
		INSERT INTO public.profiles (id, fullname)
		VALUES ($1, $2)
		RETURNING id, fullname, created_at;
	`
	row = tx.QueryRowxContext(ctx, profileQuery, profile.ID, profile.Fullname)
	if err = row.StructScan(profile); err != nil {
		return nil, nil, err
	}

	// Commit if OK
	if err = tx.Commit(); err != nil {
		return nil, nil, err
	}

	return user, profile, nil
}

func (r *RepoAuth) GetByEmail(ctx context.Context, email string) (*models.UserAuth, error) {
	var user models.UserAuth

	query := `
		SELECT 
			u.id,
			u.email,
			u.password,
			r.role AS role_name,
			u.created_at,
			u.is_verified
		FROM 
			users u
		JOIN 
			roles r ON u.role_id = r.id
		WHERE 
			u.email = $1 AND u.is_deleted = FALSE
	`

	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // no rows jadi nil, bukan error
		}
		return nil, fmt.Errorf("GetByEmail query error: %w", err)
	}

	return &user, nil
}
