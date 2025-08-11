package main

import (
	"log"
	"net/http"

	middleware "github.com/irsy4drr01/coffeeshop_be_go/internal/middlewares"
	"github.com/irsy4drr01/coffeeshop_be_go/internal/routes"
	"github.com/irsy4drr01/coffeeshop_be_go/pkg"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := pkg.Posql()

	router := routes.New(db)

	router.Static("/public", "./public")

	router.Use(middleware.CORSMiddleware())

	server := pkg.Server(router)

	log.Printf("Server is running at http://%s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v\n", err)
	}
}
