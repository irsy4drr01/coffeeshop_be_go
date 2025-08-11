package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID          int     `db:"id" json:"id,omitempty" valid:"-"`
	Uuid        string  `db:"uuid" json:"uuid,omitempty" valid:"uuid~Uuid must be a valid UUID format"`
	ProductName string  `db:"product_name" json:"product_name,omitempty" valid:"stringlength(3|100)~ProductName length must be between 3 and 100 characters"`
	Price       int     `db:"price" json:"price,omitempty" valid:"type(int)~Price must be a interger"`
	Category    string  `db:"category" json:"category,omitempty" valid:"type(string)~Price must be a string"`
	Image       string  `db:"image" json:"image" valid:"-"`
	Description *string `db:"description,omitempty" json:"description,omitempty" valid:"stringlength(0|100)~Description length max 100 characters"`
	CreatedAt   string  `db:"created_at,omitempty" json:"created_at,omitempty" valid:"-"`
	UpdatedAt   *string `db:"updated_at,omitempty" json:"updated_at,omitempty" valid:"-"`
	IsDeleted   bool    `db:"is_deleted,omitempty" json:"is_deleted,omitempty" valid:"-"`
}

type DeleteProduct struct {
	ID          int    `db:"id" json:"id,omitempty" valid:"-"`
	Uuid        string `db:"uuid" json:"uuid,omitempty" valid:"uuid~Uuid must be a valid UUID format"`
	ProductName string `db:"product_name" json:"product_name,omitempty" valid:"stringlength(3|100)~ProductName length must be between 3 and 100 characters"`
	IsDeleted   bool   `db:"is_deleted,omitempty" json:"is_deleted,omitempty" valid:"-"`
}

type Products []Product

// -----------------------------------------------------
type ProductDB struct {
	ProductID    string          `db:"id"`
	ProductName  string          `db:"product_name"`
	CategoryID   int             `db:"category_id"`
	Price        decimal.Decimal `db:"price"`
	Description  string          `db:"description"`
	TotalSold    int             `db:"total_sold"`
	TotalLikes   int             `db:"total_likes"`
	IsDeleted    bool            `db:"is_deleted"`
	CreatedAt    time.Time       `db:"created_at"`
	UpdatedAt    *time.Time      `db:"updated_at"`
	CategoryName string          `db:"category_name"`
	ProductImg   string          `db:"product_img"`
	IsDiscount   bool            `db:"is_discount"`
	DiscountRate decimal.Decimal `db:"discount_rate"`
}

type ProductResponse struct {
	ID           string `json:"id"`
	ProductName  string `json:"product_name"`
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ProductImg   string `json:"product_img"` // hanya satu gambar utama
	Description  string `json:"description"`
	TotalSold    int    `json:"total_sold"`
	TotalLike    int    `json:"total_like"`
	Price        string `json:"price"`
	FinalPrice   string `json:"final_price"`
	IsDiscount   bool   `json:"is_discount"`
	IsDeleted    bool   `json:"is_deleted"`
}

type ProductsResponse []ProductResponse

type ProductQueryParams struct {
	SearchProductName string                `form:"search_product_name"`
	Category          []CategoryQueryParams `form:"category"`
	Discount          bool                  `form:"discount"`
	SortBy            string                `form:"sort_by"`
	MinPrice          string                `form:"min_price"`
	MaxPrice          string                `form:"max_price"`
	Page              int                   `form:"page"`
	Limit             int                   `form:"limit"`
}

type CategoryQueryParams struct {
	Category1 string `form:"category1"`
	Category2 string `form:"category2"`
	Category3 string `form:"category3"`
}

// untuk SortBy
// newest -> product terbaru
// oldest -> product terlama
// asc -> urutan nama product dari a - z
// desc -> urutan nama product dari z - a
// cheapst -> product termurah
// cheapest -> product termahal
// most_liked-> by likes

type ProductDetailsResponse struct {
	ProductID   string             `json:"product_id"`
	ProductName string             `json:"product_name"`
	ProductImgs ProductImgResponse `json:"product_imgs"`
	Price       string             `json:"price"`
	FinalPrice  string             `json:"final_price"`
	TotalSold   int                `json:"total_sold"`
	TotalLike   int                `json:"total_like"`
	Description string             `json:"description"`
	IsDiscount  bool               `json:"is_discount"`
	DataSize    Sizes              `json:"data_size"`
}

type ProductImgResponse struct {
	Img1 string `json:"img_1"`
	Img2 string `json:"img_2"`
	Img3 string `json:"img_3"`
}

type Size struct {
	SizeID   int    `db:"size_id" json:"size_id"`
	SizeName string `db:"name" json:"size_name"`
}

type Sizes []Size

type ProductImgSlot struct {
	SlotNumber int    `db:"slot_number"`
	ImageFile  string `db:"image_file"`
}
