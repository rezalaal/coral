// cmd/seed/main.go
package main

import (
	"fmt"
	"log"

	"github.com/rezalaal/coral/internal/db"
	"github.com/rezalaal/coral/internal/models"
	"github.com/rezalaal/coral/internal/repository/postgres"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("خطا در اتصال به دیتابیس:", err)
	}

	userRepo := postgres.NewUserPG(database)

	// کاربران اولیه
	users := []models.User{
		{Name: "علی", Mobile: "09121234567", PasswordHash: "hashed-password-1"},
		{Name: "زهرا", Mobile: "09391234567", PasswordHash: "hashed-password-2"},
	}
	for _, u := range users {
		err := userRepo.Create(&u)
		if err != nil {
			fmt.Println("❌ خطا در ساخت کاربر:", err)
		} else {
			fmt.Println("✅ کاربر ساخته شد:", u.Name)
		}
	}

}
