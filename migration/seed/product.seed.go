package seed

import (
	"log"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/models"
	"github.com/jmoiron/sqlx"
)

func SeedProducts(db *sqlx.DB) {
	products := models.Products{
		{
			ProductName: "Americano Coffee",
			Price:       15000,
			Category:    "Coffee",
			Description: stringDesPtr("Kopi berbasis espresso dengan air panas."),
		},
		{
			ProductName: "Double Espresso Coffee",
			Price:       22000,
			Category:    "Coffee",
			Description: stringDesPtr("Kopi kuat dan pekat yang dibuat dengan memaksa air panas melalui biji kopi halus."),
		},
		{
			ProductName: "Nasi Goreng Spesial",
			Price:       20000,
			Category:    "Food",
			Description: stringDesPtr("Nasi goreng spesial dengan campuran sayuran, ayam, dan telur."),
		},
		{
			ProductName: "Chicken Sandwich",
			Price:       20000,
			Category:    "Food",
			Description: stringDesPtr("Sandwich ayam panggang dengan selada, tomat, dan mayones di roti gandum."),
		},
		{
			ProductName: "Kentang Goreng",
			Price:       10000,
			Category:    "Snack",
			Description: stringDesPtr("Irisan kentang goreng yang renyah dengan bumbu garam."),
		},
		{
			ProductName: "Pempek Bakar",
			Price:       10000,
			Category:    "Snack",
			Description: stringDesPtr("Pempek panggang, hidangan tradisional Indonesia, disajikan dengan saus cuka pedas."),
		},
		{
			ProductName: "Air Mineral",
			Price:       5000,
			Category:    "Non-Coffee",
			Description: stringDesPtr("Air mineral kemasan, 500 ml."),
		},
		{
			ProductName: "Lemon Iced Tea",
			Price:       13000,
			Category:    "Non-Coffee",
			Description: stringDesPtr("Es teh yang segar dengan sedikit perasan lemon."),
		},
		{
			ProductName: "Milk Coffee",
			Price:       28000,
			Category:    "Coffee",
			Description: stringDesPtr("Kopi susu yang lembut dan nikmat, dibuat dengan campuran espresso dan susu segar, menciptakan minuman yang kaya rasa dan creamy."),
		},
		{
			ProductName: "Vanilla Latte",
			Price:       29000,
			Category:    "Coffee",
			Description: stringDesPtr("Smooth vanilla-infused espresso with creamy steamed milk, perfect for a delightful pick me up."),
		},
		{
			ProductName: "Hazelnut Latte Coffee",
			Price:       20000,
			Category:    "Coffee",
			Description: stringDesPtr("Smooth espresso with rich hazelnut flavor and creamy steamed milk, a comforting treat."),
		},
		{
			ProductName: "Hazelnut Chocolate",
			Price:       28000,
			Category:    "Non-Coffee",
			Description: stringDesPtr("Indulgent hazelnut-infused chocolate delight, perfect for a rich, creamy treat."),
		},
		{
			ProductName: "Palm Milk Coffee",
			Price:       27000,
			Category:    "Coffee",
			Description: stringDesPtr("Two shots of espresso with pure fresh milk and palm sugar creates a harmonious flavor."),
		},
		{
			ProductName: "Sweet Iced Tea",
			Price:       18000,
			Category:    "Non-Coffee",
			Description: stringDesPtr("Fresh tea with ice and separate sugar, mix your own sugar and enjoy its freshness."),
		},
		{
			ProductName: "Test 4",
			Price:       18000,
			Category:    "Non-Coffee",
			Description: stringDesPtr("Description test."),
		},
	}

	query := `INSERT INTO public.product (product_name, price, category, description) 
		VALUES (:product_name, :price, :category, :description)`

	for _, product := range products {
		_, err := db.NamedExec(query, product)
		if err != nil {
			log.Printf("Failed to seed product: %s. Error: %v", product.ProductName, err)
		} else {
			log.Printf("Successfully seeded product: %s", product.ProductName)
		}
	}
}

func stringDesPtr(s string) *string {
	return &s
}
