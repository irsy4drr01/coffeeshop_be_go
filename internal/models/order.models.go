package models

import "github.com/shopspring/decimal"

// akan insert ke tabel orders
type CreateOrderDB struct {
	OrderID           string          `db:"order_id"`            // column id di table orders, generate order_id dengan format TRXYYYYMMDDHHMMSSXXXXXXXX (x di belakang adalah uuid yang dipotong 8 karakter)
	UserID            string          `db:"user_id"`             // column user_id di table orders, sepertinya didapat dari token ctx.Get("")
	FullName          string          `db:"fullname"`            // column fullname di table orders, sebagai perekam data agar tidak berubah di riwayat
	Address           string          `db:"address"`             // column address di table orders, sebagai perekam data agar tidak berubah di riwayat
	Phone             string          `db:"phone"`               // column phone di table orders, sebagai perekam data agar tidak berubah di riwayat
	PaymentMethod     string          `db:"payment_method"`      // column payment_method di table orders, sebagai perekam data agar tidak berubah di riwayat
	TotalPurchase     string          `db:"total_purchase"`      // column total_purchase di table orders, sebagai perekam data agar tidak berubah di riwayat
	DeliveryMethod    string          `db:"delivery_method"`     // column delivery_method di table orders, sebagai perekam data agar tidak berubah di riwayat
	DeliveryMethodFee decimal.Decimal `db:"delivery_method_fee"` // column delivery_method_fee di table orders, sebagai perekam data agar tidak berubah di riwayat
	Tax               float64         `db:"tax"`                 // column tax di table orders, sebagai perekam data agar tidak berubah di riwayat
	StatusID          int             `db:"status_id"`           // column status_id di table orders,
	TotalAmount       decimal.Decimal `db:"total_amount"`        // column total_amount di table orders, sebagai perekam data agar tidak berubah di riwayat
}

// akan insert ke tabel order_details
type OrderDetailsDB struct {
	OrderID      string          `db:"order_id"`      // insert sama dengan order_id di tabel orders dengan column id
	ProductID    string          `db:"product_id"`    // insert id product sesuai product yang dipilih user
	Qty          int             `db:"qty"`           // insert quantity
	SizeID       int             `db:"size_id"`       // insert size id
	IsIced       bool            `db:"is_iced"`       // insert apakah ice (true/false)
	IsDiscount   bool            `db:"is_discount"`   // insert apakah discount (true/false)
	DiscountName string          `db:"discount_name"` // column discount_name di table order_details, sebagai perekam data agar tidak berubah di riwayat
	BasePrice    decimal.Decimal `db:"base_price"`    // column base_price di table order_details, sebagai perekam data agar tidak berubah di riwayat, didapat dari price di tabel products
	FinalPrice   decimal.Decimal `db:"final_price"`   // column final_price di table order_details, sebagai perekam data agar tidak berubah di riwayat, didapat dari final_price di tabel products
	SubTotal     decimal.Decimal `db:"sub_total"`     // column sub_total di table order_details
}

// logic order dan order_details
// sub total = quantity * (harga asli + (harga asli * persen penambahan size) + penambahan jika is_iced true (diambil dari harga Ice Cube, bukan extra Ice Cube))
// total purchase = sub total product 1 + sub total product 2 + .........
// total amount = total purchase + delivery fee (jika menggunakan delivery) + (tax * total purchase)

type ProductWithSizeAndDiscount struct {
	ProductID      string          `db:"product_id"`
	ProductName    string          `db:"product_name"`
	BasePrice      decimal.Decimal `db:"price"`
	SizeID         int             `db:"size_id"`
	SizeName       string          `db:"size_name"`
	Additional     float64         `db:"additional_price"`
	CategoryID     int             `db:"category_id"`
	DiscountID     *int            `db:"discount_id"`
	DiscountName   *string         `db:"discount_name"`
	DiscountValue  *float64        `db:"discount_value"`
	DiscountActive *bool           `db:"is_actived"`
	DiscountExpiry *string         `db:"expired"`
}

type DeliveryMethod struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Fee  int    `db:"fee"`
}

type PaymentMethod struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Tax struct {
	ID       int     `db:"id"`
	TaxValue float64 `db:"tax_value"`
}

type CreateOrderRequest struct {
	UserID           string             `json:"-"`
	Phone            string             `json:"phone"`
	Fullname         string             `json:"fullname"`
	Address          string             `json:"address"`
	DeliveryMethodID int                `json:"delivery_method_id"`
	PaymentMethodID  int                `json:"payment_method_id"`
	Items            []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID string `json:"product_id"`
	Qty       int    `json:"qty"`
	SizeID    int    `json:"size_id"`
	IsIced    bool   `json:"is_iced"`
}

type CreateOrderRepoResult struct {
	OrderID        string
	Phone          string
	Items          []OrderDetailsResponse
	DeliveryMethod DeliveryMethod
	PaymentMethod  PaymentMethod
	TaxAmount      decimal.Decimal
	TotalPurchase  decimal.Decimal
	TotalAmount    decimal.Decimal
}

type CreateOrderResponse struct {
	OrderID          string                 `json:"order_id"`
	UserID           string                 `json:"user_id"`
	Fullname         string                 `json:"fullname"`
	Address          string                 `json:"address"`
	Phone            string                 `json:"phone"`
	DeliveryMethodID int                    `json:"delivery_method_id"`
	DeliveryMethod   string                 `json:"delivery_method"`
	PaymentMethodID  int                    `json:"payment_method_id"`
	PaymentMethod    string                 `json:"payment_method"`
	Items            []OrderDetailsResponse `json:"items"`
	TotalPurchase    string                 `json:"total_purchase"`
	DeliveryFee      string                 `json:"delivery_fee"`
	TaxAmount        string                 `json:"tax_amount"`
	TotalAmount      string                 `json:"total_amount"`
}

type OrderDetailsResponse struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Qty         int    `json:"qty"`
	Size        string `json:"size"`
	IsIced      bool   `json:"is_iced"`
	BasePrice   string `json:"base_price"`
	FinalPrice  string `json:"final_price"`
	SubTotal    string `json:"sub_total"`
}
