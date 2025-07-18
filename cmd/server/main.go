package main

import (
	"log"
	"net/http"

	"github.com/rezalaal/coral/internal/db"
	"github.com/rezalaal/coral/internal/repository/postgres"
	"github.com/rezalaal/coral/internal/router"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	userRepo := postgres.NewUserRepository(dbConn)
	restaurantRepo := postgres.NewRestaurantRepository(dbConn)

	r := router.NewRouter(userRepo, restaurantRepo)

	log.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
