package main

import (
	"context"
	"log"

	"github.com/irsy4drr01/coffeeshop_be_go/migration/seed"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()

	db := pkg.Posql()
	defer db.Close()

	// log.Println("Starting roles seeding...")
	// seed.SeedRoles(ctx, db)

	// log.Println("Starting users seeding...")
	// seed.SeedUsers(ctx, db)

	// log.Println("Starting profiles seeding...")
	// seed.SeedProfiles(ctx, db)

	log.Println("Seeding roles, users and profiles (transactional)...")
	if err := seed.SeedRoleSAndUsersAndProfles(ctx, db); err != nil {
		log.Fatalf("Error seeding roles, users, profiles: %v", err)
	}

	log.Println("Seeding categories...")
	if err := seed.SeedCategories(ctx, db); err != nil {
		log.Fatalf("Error seeding categories: %v", err)
	}

	log.Println("Seeding sizes...")
	if err := seed.SeedSizes(ctx, db); err != nil {
		log.Fatalf("Error seeding sizes: %v", err)
	}

	log.Println("Seeding products...")
	if err := seed.SeedProducts(ctx, db); err != nil {
		log.Fatalf("Error seeding products: %v", err)
	}

	log.Println("Seeding product_stocks...")
	if err := seed.SeedProductStocks(ctx, db); err != nil {
		log.Fatalf("Error seeding product_stocks: %v", err)
	}

	log.Println("Seeding product_images...")
	if err := seed.SeedProductImages(ctx, db); err != nil {
		log.Fatalf("Error seeding product_images: %v", err)
	}

	log.Println("Seeding discounts...")
	if err := seed.SeedDiscounts(ctx, db); err != nil {
		log.Fatalf("Error seeding discounts: %v", err)
	}

	log.Println("Seeding discount_products...")
	if err := seed.SeedDiscountProducts(ctx, db); err != nil {
		log.Fatalf("Error seeding discount_products: %v", err)
	}

	log.Println("Seeding payment_methods...")
	if err := seed.SeedPaymentMethods(ctx, db); err != nil {
		log.Fatalf("Error seeding payment_methods: %v", err)
	}

	log.Println("Seeding delivery_methods...")
	if err := seed.SeedDeliveryMethods(ctx, db); err != nil {
		log.Fatalf("Error seeding delivery_methods: %v", err)
	}

	log.Println("Seeding statuses...")
	if err := seed.SeedStatuses(ctx, db); err != nil {
		log.Fatalf("Error seeding statuses: %v", err)
	}

	log.Println("Seeding tax...")
	if err := seed.SeedTax(ctx, db); err != nil {
		log.Fatalf("Error seeding tax: %v", err)
	}

	// log.Println("Seeding orders...")
	// if err := seed.SeedTax(ctx, db); err != nil {
	// 	log.Fatalf("Error seeding orders: %v", err)
	// }

	// log.Println("Seeding order_details...")
	// if err := seed.SeedTax(ctx, db); err != nil {
	// 	log.Fatalf("Error seeding order_details: %v", err)
	// }

	// log.Println("Seeding product_likes...")
	// if err := seed.SeedTax(ctx, db); err != nil {
	// 	log.Fatalf("Error seeding product_likes: %v", err)
	// }

	log.Println("All Seeding completed successfully.")
}
