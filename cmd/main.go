package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                            // domain yang diperbolehkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // method yang diperbolehkan
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // durasi cache hasil preflight request
	}))

	// router.Use(cors.Default())
	// Default CORS if use cors.Default()
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	ExposeHeaders:    nil,
	// 	AllowCredentials: false,
	// 	MaxAge: 0,
	// }))

	server := pkg.Server(router)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
