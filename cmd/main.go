package main

import (
	"log"

	"github.com/irsy4drr01/coffeeshop_be_go/internal/routes"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := pkg.Posql()
	if err != nil {
		log.Fatal(err)
	}

	router := routes.New(db)
	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
