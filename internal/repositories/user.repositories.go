package repositories

import (
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type UserRepoInterface interface {
	CreateUser(data *models.User) (string, *models.User, error)
	GetAllUser(searchUserName string, sort string, page int, limit int) (*models.Users, error)
	GetOneUser(uuid string) (*models.User, error)
	UpdateUser(uuid string, body map[string]any) (string, *models.User, error)
	DeleteUser(uuid string) (string, *models.User, error)
}

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) UserRepoInterface {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, *models.User, error) {
	query := `
		INSERT INTO public.users (		
			username,
			email,
			password
		)
		VALUES (:username, :email, :password)
		RETURNING uuid, username, email, password, created_at;
	`

	var user models.User
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

func (r *RepoUser) GetAllUser(searchUserName string, sort string, page int, limit int) (*models.Users, error) {
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
		WHERE username ILIKE '%' || $1 || '%' AND is_deleted = false
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

	// Tambahkan pagination
	offset := (page - 1) * limit
	query += " LIMIT $2 OFFSET $3;"

	data := models.Users{}
	if err := r.Select(&data, query, searchUserName, limit, offset); err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *RepoUser) GetOneUser(uuid string) (*models.User, error) {
	query := `
		SELECT
			uuid,
			username,
			email,
			password,
			created_at
		FROM public.users
		WHERE uuid = $1 AND is_deleted = false;
	`

	var userDetail models.User
	if err := r.Get(&userDetail, query, uuid); err != nil {
		return nil, err
	}
	return &userDetail, nil
}

func (r *RepoUser) UpdateUser(uuid string, body map[string]any) (string, *models.User, error) {
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

	var user models.User
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

func (r *RepoUser) DeleteUser(uuid string) (string, *models.User, error) {
	query := `
		UPDATE public.users
		SET
			is_deleted = true
		WHERE uuid = $1
		RETURNING username, email, uuid, is_deleted;
	`

	var user models.User
	if err := r.Get(&user, query, uuid); err != nil {
		return "", nil, err
	}
	return "User deleted successfully.", &user, nil
}
