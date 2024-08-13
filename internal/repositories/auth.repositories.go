package repositories

import (
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepoInterface interface {
	CreateUser(data *models.User) (string, *models.User, error)
	GetByEmail(email string) (*models.User, error)
}

type RepoAuth struct {
	*sqlx.DB
}

func NewAuth(db *sqlx.DB) *RepoAuth {
	return &RepoAuth{db}
}

func (r *RepoAuth) CreateUser(data *models.User) (string, *models.User, error) {
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

func (r *RepoAuth) GetByEmail(email string) (*models.User, error) {
	result := models.User{}
	query := `SELECT uuid, username, email, password FROM public.users WHERE email=$1 and is_deleted = false`
	err := r.Get(&result, query, email)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
