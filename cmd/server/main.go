package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"coral/internal/router"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not set in .env")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	// DI Container with secrets
	c := router.NewContainer(db)

	// HTTP Router
	r := router.NewRouter(c)

	// Start Server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
