package config

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

// Config ساختاری برای نگهداری تنظیمات پروژه
type Config struct {
	DatabaseURL      string
	KavenegarAPIKey  string
	KavenegarTemplate string
}

// Load بارگذاری تنظیمات از فایل .env
func Load() (*Config, error) {
	// دریافت مسیر کاری فعلی
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("خطا در دریافت مسیر کاری: %v", err)
	}

	// پیدا کردن مسیر ریشه پروژه که حاوی فایل go.mod است
	for {
		// بررسی وجود فایل go.mod در مسیر فعلی
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			// وقتی go.mod پیدا شد، مسیر ریشه پروژه است
			break
		}

		// اگر go.mod پیدا نشد، به پوشه بالاتر برویم
		currentDir = filepath.Dir(currentDir)
	}

	// ساخت مسیر به فایل .env از ریشه پروژه
	rootPath := filepath.Join(currentDir, ".env")

	// چاپ مسیر برای دیباگ
	fmt.Println("Loading .env from path:", rootPath)

	// بارگذاری فایل .env
	err = godotenv.Load(rootPath)
	if err != nil {
		return nil, fmt.Errorf("خطا در بارگذاری فایل .env از مسیر %s: %v", rootPath, err)
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
