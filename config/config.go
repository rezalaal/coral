package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

// Config ساختاری برای نگهداری تنظیمات پروژه
type Config struct {
	DatabaseURL      string
	KavenegarAPIKey  string
	KavenegarTemplate string
}

// Load بارگذاری تنظیمات از فایل .env
func Load() (*Config, error) {
	// تعیین مسیر دقیق فایل .env
	rootPath, _ := filepath.Abs(filepath.Join("..", "..", ".env"))
	err := godotenv.Load(rootPath)
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
