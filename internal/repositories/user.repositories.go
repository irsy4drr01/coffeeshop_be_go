package repositories

import (
	"context"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepoInterface interface {
	// GetAllUser(searchUserName string, sort string, page int, limit int) (*models.Users, error)
	GetOneUser(ctx context.Context, id string) (*models.UserDB, error)
	// UpdateUser(ctx context.Context, id string, .........) error
	// DeleteUser(uuid string) (*models.User, error)
}

type RepoUser struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db: db}
}

// func (r *RepoUser) GetAllUser(searchUserName string, sort string, page int, limit int) (*models.Users, error) {
// 	query := `
// 		SELECT
// 			uuid,
// 			username,
// 			password,
// 			email,
// 			image,
// 			created_at,
// 			updated_at,
// 			is_deleted
// 		FROM public.users
// 		WHERE username ILIKE '%' || $1 || '%' AND is_deleted = false
// 		ORDER BY `

// 	// Tambahkan logika sort berdasarkan parameter
// 	switch sort {
// 	case "ASC":
// 		query += "username ASC"
// 	case "DESC":
// 		query += "username DESC"
// 	case "oldest":
// 		query += "created_at ASC"
// 	case "newest":
// 		query += "created_at DESC"
// 	default:
// 		query += "created_at DESC" // default sort
// 	}

// 	// Tambahkan pagination
// 	offset := (page - 1) * limit
// 	query += " LIMIT $2 OFFSET $3;"

// 	data := models.Users{}
// 	if err := r.Select(&data, query, searchUserName, limit, offset); err != nil {
// 		return nil, err
// 	}
// 	return &data, nil
// }

// ------------------------------------------------

func (r *RepoUser) GetOneUser(ctx context.Context, id string) (*models.UserDB, error) {
	var user models.UserDB
	query := `
		SELECT 
			u.id,
			u.email,
			r.role,
			p.fullname,
			p.phone,
			p.address,
			COALESCE(p.image, 'avatar_default.webp') AS image,
			u.is_verified,
			u.is_deleted,
			u.created_at,
			u.updated_at
		FROM users u
		JOIN roles r ON u.role_id = r.id
		JOIN profiles p ON p.id = u.id
		WHERE u.id = $1 AND u.is_deleted = false
	`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ------------------------------------------------

// func (r *RepoUser) UpdateUser(ctx context.Context, id string, reqBodyDB models.UserDB) error {

// }

// ------------------------------------------------

// func (r *RepoUser) UpdateUser(uuid string, body map[string]any) (*models.User, error) {
// 	query := `UPDATE users SET `
// 	params := map[string]interface{}{}

// 	if username, exists := body["username"]; exists {
// 		query += "username = :username, "
// 		params["username"] = username
// 	}
// 	if email, exists := body["email"]; exists {
// 		query += "email = :email, "
// 		params["email"] = email
// 	}
// 	if password, exists := body["password"]; exists {
// 		query += "password = :password, "
// 		params["password"] = password
// 	}
// 	if image, exists := body["image"]; exists {
// 		query += "image = :image, "
// 		params["image"] = image
// 	}

// 	query += "updated_at = NOW() WHERE uuid = :uuid RETURNING username, email, password, uuid, image, updated_at"
// 	params["uuid"] = uuid

// 	var user models.User
// 	stmt, args, err := sqlx.Named(query, params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	stmt = r.Rebind(stmt) // Rebind the statement according to the driver
// 	if err := r.QueryRowx(stmt, args...).StructScan(&user); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// ------------------------------------------------

// func (r *RepoUser) DeleteUser(uuid string) (*models.User, error) {
// 	query := `
// 		UPDATE public.users
// 		SET
// 			is_deleted = true
// 		WHERE uuid = $1
// 		RETURNING username, email, uuid, is_deleted;
// 	`

// 	var user models.User
// 	if err := r.Get(&user, query, uuid); err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }
