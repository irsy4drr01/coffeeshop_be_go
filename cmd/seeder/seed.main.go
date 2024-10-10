package main

import (
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

	db, err := pkg.Posql()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// log.Println("Starting products seeding...")
	// seed.SeedProducts(db)

	log.Println("Starting users seeding...")
	seed.SeedUsers(db)

	log.Println("Seeding completed successfully.")
}
