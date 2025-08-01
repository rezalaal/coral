// cmd/seed/main.go
package main

import (
	"fmt"
	"log"

	"coral/internal/db"
	"coral/internal/domain/user"
	"coral/internal/infrastructure/postgres"
)


func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("خطا در اتصال به دیتابیس:", err)
	}

	userRepo := postgres.NewUserPG(database)

	// کاربران اولیه
	users := []user.User{
		{Name: "علی", Mobile: "09121234567", Password: "hashed-password-1"},
		{Name: "زهرا", Mobile: "09391234567", Password: "hashed-password-2"},
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
