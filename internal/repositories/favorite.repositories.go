package repositories

import (
	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

type RepoFavorite struct {
	*sqlx.DB
}

func NewFavorite(db *sqlx.DB) *RepoFavorite {
	return &RepoFavorite{db}
}

func (r *RepoFavorite) AddFavorite(userID, productID int) (string, *models.FavoriteResponse, error) {
	insertQuery := `
        INSERT INTO public.favorite (
            user_id,
            product_id
        )
        VALUES ($1, $2)
        ON CONFLICT (user_id, product_id) DO NOTHING;
    `

	_, err := r.Exec(insertQuery, userID, productID)
	if err != nil {
		return "", nil, err
	}

	selectQuery := `
        SELECT
            f.user_id,
            u.username,
            f.product_id,
            p.product_name
        FROM public.favorite f
        JOIN public.users u ON f.user_id = u.id
        JOIN public.product p ON f.product_id = p.id
        WHERE f.user_id = $1 AND f.product_id = $2;
    `

	var favorite models.FavoriteResponse
	if err := r.Get(&favorite, selectQuery, userID, productID); err != nil {
		return "", nil, err
	}

	return "Product added to favorites successfully.", &favorite, nil
}

func (r *RepoFavorite) RemoveFavorite(userID, productID int) (string, error) {
	query := `
		DELETE FROM public.favorite
		WHERE user_id = $1 AND product_id = $2;
	`

	if _, err := r.Exec(query, userID, productID); err != nil {
		return "", err
	}
	return "Product removed from favorites successfully.", nil
}

func (r *RepoFavorite) GetFavorites(userID int) (*models.Favorites, error) {
	query := `
		SELECT
			f.user_id,
			u.username,
			f.product_id,
			p.product_name
		FROM public.favorite f
		JOIN public.users u ON f.user_id = u.id
		JOIN public.product p ON f.product_id = p.id
		WHERE f.user_id = $1;
	`

	var favorites models.Favorites
	if err := r.Select(&favorites, query, userID); err != nil {
		return nil, err
	}
	return &favorites, nil
}
