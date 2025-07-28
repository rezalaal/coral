package main

import (
    "log"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/rezalaal/coral/internal/db"
    "github.com/rezalaal/coral/internal/user/repository/postgres"
    "github.com/rezalaal/coral/internal/router"
)

func main() {
    // بارگذاری فایل .env
    err := godotenv.Load()
    if err != nil {
        log.Println(".env not found")
    }

    // اتصال به دیتابیس
    dbConn, err := db.Connect()
    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    // ایجاد repository برای user
    userRepo := postgres.NewUserPG(dbConn)

    // ایجاد و راه‌اندازی روت‌ها
    r := router.NewRouter(dbConn, userRepo, nil)

    // راه‌اندازی سرور
    log.Println("Server running on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}
