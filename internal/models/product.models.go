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

// type Product struct {
// 	ID          int     `db:"id" json:"id"`
// 	Uuid        string  `db:"uuid" json:"uuid"`
// 	ProductName string  `db:"product_name" json:"product_name"`
// 	Price       int     `db:"price" json:"price"`
// 	Category    string  `db:"category" json:"category"`
// 	Description *string `db:"description" json:"description"`
// 	CreatedAt   string  `db:"created_at" json:"created_at"`
// 	UpdatedAt   *string `db:"updated_at" json:"updated_at"`
// 	IsDeleted   bool    `db:"is_deleted" json:"is_deleted"`
// }

// type ProductDetail struct {
// 	ID          int     `db:"id" json:"id"`
// 	Uuid        string  `db:"uuid" json:"uuid"`
// 	ProductName string  `db:"product_name" json:"product_name"`
// 	Price       int     `db:"price" json:"price"`
// 	Category    string  `db:"category" json:"category"`
// 	Description *string `db:"description" json:"description"`
// }

// type CreateProductResponse struct {
// 	ID          int     `db:"id" json:"id"`
// 	Uuid        string  `db:"uuid" json:"uuid"`
// 	ProductName string  `db:"product_name" json:"product_name"`
// 	Price       int     `db:"price" json:"price"`
// 	Category    string  `db:"category" json:"category"`
// 	Description *string `db:"description" json:"description"`
// 	CreatedAt   string  `db:"created_at" json:"created_at"`
// }

// type UpdateProductResponse struct {
// 	ID          int     `db:"id" json:"id"`
// 	Uuid        string  `db:"uuid" json:"uuid"`
// 	ProductName string  `db:"product_name" json:"product_name"`
// 	Price       int     `db:"price" json:"price"`
// 	Category    string  `db:"category" json:"category"`
// 	Description *string `db:"description" json:"description"`
// 	UpdatedAt   *string `db:"updated_at" json:"updated_at"`
// }

// type DeleteProductResponse struct {
// 	ID          int    `db:"id" json:"id"`
// 	Uuid        string `db:"uuid" json:"uuid"`
// 	ProductName string `db:"product_name" json:"product_name"`
// 	Category    string `db:"category" json:"category"`
// 	IsDeleted   bool   `db:"is_deleted" json:"is_deleted"`
// }

type Products []Product
