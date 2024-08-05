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

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `
		INSERT INTO public.users (		
			username,
			email,
			password
		)
		VALUES (:username, :email, :password);
	`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "User has been created", nil
}

func (r *RepoUser) GetAllUser() (*models.Users, error) {
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
		order by created_at DESC;
	`
	data := models.Users{}

	if err := r.Select(&data, query); err != nil {
		return nil, err
	}
	return &data, nil
}
