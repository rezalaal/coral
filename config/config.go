package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config ساختاری برای نگهداری تنظیمات پروژه
type Config struct {
	DatabaseURL      string
	KavenegarAPIKey  string
	KavenegarTemplate string
}

// Load بارگذاری تنظیمات از فایل .env
func Load() (*Config, error) {
	// بارگذاری فایل .env
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("خطا در بارگذاری فایل .env: %v", err)
	}

	// دریافت تنظیمات از محیط
	config := &Config{
		DatabaseURL:      os.Getenv("DATABASE_URL"),
		KavenegarAPIKey:  os.Getenv("KAVENEGAR_API_KEY"),
		KavenegarTemplate: os.Getenv("KAVENEGAR_TEMPLATE"),
	}

	// بررسی اینکه همه تنظیمات به درستی بارگذاری شده‌اند
	if config.DatabaseURL == "" || config.KavenegarAPIKey == "" || config.KavenegarTemplate == "" {
		return nil, fmt.Errorf("تنظیمات ضروری از فایل .env یافت نشد")
	}

	return config, nil
}
