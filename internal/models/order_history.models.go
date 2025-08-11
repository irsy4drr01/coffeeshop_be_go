package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// type OrderHistoryDB struct {
// 	ID                string
// 	UserID            string
// 	Image             string
// 	FullName          string
// 	Address           string
// 	Phone             string
// 	PaymentMethod     string
// 	DeliveryMethod    string
// 	DeliveryMethodFee decimal.Decimal
// 	CreatedAt         time.Time
// 	TotalAmount       decimal.Decimal
// 	Status            string
// }

// ------order history------

type OrderHistoryDB struct {
	ID          string          `db:"order_id"`
	Image       string          `db:"image"`
	Status      string          `db:"status"`
	CreatedAt   time.Time       `db:"created_at"`
	TotalAmount decimal.Decimal `db:"total_amount"`
}

type OrderHistoriesDB []OrderHistoryDB

// type OrderHistoryRes struct {
// 	ID                string
// 	Image             string
// 	FullName          string
// 	Address           string
// 	Phone             string
// 	PaymentMethod     string
// 	DeliveryMethod    string
// 	DeliveryMethodFee string
// 	CreatedAt         string
// 	TotalAmount       string
// 	Status            string
// }

type OrderHistoryRes struct {
	ID          string `json:"order_id"`
	Image       string `json:"image"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	TotalAmount string `json:"total_amount"`
}

type OrderHistoriesRes []OrderHistoryRes

// ------order history detail------

type OrderHistoryDetailDB struct {
	ID                string          `db:"order_id"`
	FullName          string          `db:"fullname"`
	Address           string          `db:"address"`
	Phone             string          `db:"phone"`
	PaymentMethod     string          `db:"payment_method"`
	Status            string          `db:"status_name"`
	TotalPurchase     decimal.Decimal `db:"total_purchase"`
	DeliveryMethod    string          `db:"delivery_method"`
	DeliveryMethodFee decimal.Decimal `db:"delivery_method_fee"`
	Tax               decimal.Decimal `db:"tax"`
	TotalAmount       decimal.Decimal `db:"total_amount"`
	CreatedAt         time.Time       `db:"created_at"`
	ProductName       string          `db:"product_name"`
	Image             string          `db:"image"`
	Qty               int             `db:"qty"`
	Size              string          `db:"size_name"`
	IsIced            bool            `db:"is_iced"`
	BasePrice         decimal.Decimal `db:"base_price"`
	FinalPrice        decimal.Decimal `db:"finale_price"`
	DiscountName      *string         `db:"discount_name"`
	Items             OrderHistoryItemsDB
}

type OrderHistoryItemDB struct {
	ProductName  string          `db:"product_name"`
	Image        string          `db:"image"`
	Qty          int             `db:"qty"`
	Size         string          `db:"size_name"`
	IsIced       bool            `db:"is_iced"`
	BasePrice    decimal.Decimal `db:"base_price"`
	FinalPrice   decimal.Decimal `db:"finale_price"`
	DiscountName string          `db:"discount_name"`
	CategoryID   int             `db:"category_id"`
}

type OrderHistoryItemsDB []OrderHistoryItemDB

type OrderHistoryDetailRes struct {
	ID                string `json:"order_id"`
	FullName          string `json:"fullname"`
	Address           string `json:"address"`
	Phone             string `json:"phone"`
	PaymentMethod     string `json:"payment_method"`
	Status            string `json:"status"`
	TotalPurchase     string `json:"total_purchase"`
	DeliveryMethod    string `json:"delivery_method"`
	DeliveryMethodFee string `json:"delivery_method_fee"`
	Tax               string `json:"tax"`
	TotalAmount       string `json:"total_amount"`
	CreatedAt         string `json:"created_at"`
	Items             OrderHistoryItemsRes
}

type OrderHistoryItemRes struct {
	ProductName  string `json:"product_name"`
	Image        string `json:"image"`
	Qty          int    `json:"qty"`
	Size         string `json:"size"`
	Temperature  string `json:"temperature"`
	BasePrice    string `json:"base_price"`
	FinalPrice   string `json:"finale_price"`
	DiscountName string `json:"discount_name"`
}

type OrderHistoryItemsRes []OrderHistoryItemRes
