package repositories

import (
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, *models.CreateUserResponse, error) {
	query := `
		INSERT INTO public.users (		
			username,
			email,
			password
		)
		VALUES (:username, :email, :password)
		RETURNING uuid, username, email, password, created_at;
	`

	var user models.CreateUserResponse
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return "", nil, err
	}

	err = stmt.Get(&user, data)
	stmt.Close() // Menutup statement setelah digunakan
	if err != nil {
		return "", nil, err
	}
	return "User created successfully.", &user, nil
}

func (r *RepoUser) GetAllUser(searchUserName string, sort string) (*models.Users, error) {
	query := `
		SELECT
			uuid,
			username,
			password,
			email,
			created_at,
			updated_at,
			is_deleted
		FROM public.users
		WHERE username ILIKE '%' || $1 || '%'
		ORDER BY `

	// Tambahkan logika sort berdasarkan parameter
	switch sort {
	case "a-z":
		query += "username ASC"
	case "z-a":
		query += "username DESC"
	case "oldest":
		query += "created_at ASC"
	case "newest":
		query += "created_at DESC"
	default:
		query += "created_at DESC" // default sort
	}

	query += ";"

	data := models.Users{}
	if err := r.Select(&data, query, searchUserName); err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoUser) UpdateUser(uuid string, body map[string]any) (string, *models.UpdateUserResponse, error) {
	query := `UPDATE users SET `
	params := map[string]interface{}{}

	if username, exists := body["username"]; exists {
		query += "username = :username, "
		params["username"] = username
	}
	if emailValue, exists := body["email"]; exists {
		query += "email = :email, "
		params["email"] = emailValue
	}
	if password, exists := body["password"]; exists {
		query += "password = :password, "
		params["password"] = password
	}

	query += "updated_at = NOW() WHERE uuid = :uuid RETURNING username, email, password, uuid, updated_at"
	params["uuid"] = uuid

	var user models.UpdateUserResponse
	stmt, args, err := sqlx.Named(query, params)
	if err != nil {
		return "", nil, err
	}

	stmt = r.Rebind(stmt) // Rebind the statement according to the driver
	if err := r.QueryRowx(stmt, args...).StructScan(&user); err != nil {
		return "", nil, err
	}
	return "User updated successfully.", &user, nil
}

func (r *RepoUser) DeleteUser(uuid string) (string, *models.DeleteUserResponse, error) {
	query := `
		UPDATE public.users
		SET
			is_deleted = true
		WHERE uuid = $1
		RETURNING username, email, uuid, is_deleted;
	`

	var user models.DeleteUserResponse
	if err := r.Get(&user, query, uuid); err != nil {
		return "", nil, err
	}
	return "User deleted successfully.", &user, nil
}
