package models

var schemaProductsFavorites = `
CREATE TABLE public.favorite (
    user_id int4 NOT NULL,
    product_id int4 NOT NULL,
    CONSTRAINT favorite_pk PRIMARY KEY (user_id, product_id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users (id) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES public.product (id) ON DELETE CASCADE
);`

var _ = schemaProductsFavorites

type Favorite struct {
	UserID      int    `db:"user_id" json:"user_id"`
	Username    string `db:"username,omitempty" json:"username,omitempty"`
	ProductID   int    `db:"product_id" json:"product_id"`
	ProductName string `db:"product_name,omitempty" json:"product_name,omitempty"`
}

type Favorites []Favorite
