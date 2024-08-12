package models

type Product struct {
	ID          int     `db:"id" json:"id,omitempty"`
	Uuid        string  `db:"uuid" json:"uuid,omitempty"`
	ProductName string  `db:"product_name" json:"product_name,omitempty"`
	Price       int     `db:"price" json:"price,omitempty"`
	Category    string  `db:"category" json:"category,omitempty"`
	Description *string `db:"description,omitempty" json:"description,omitempty"`
	CreatedAt   string  `db:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt   *string `db:"updated_at,omitempty" json:"updated_at,omitempty"`
	IsDeleted   bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty"`
}

type Products []Product
