package models

type Favorite struct {
	UserID      int    `db:"user_id" json:"user_id"`
	Username    string `db:"username,omitempty" json:"username,omitempty"`
	ProductID   int    `db:"product_id" json:"product_id"`
	ProductName string `db:"product_name,omitempty" json:"product_name,omitempty"`
}

type Favorites []Favorite
