package config_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/rezalaal/coral/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig_Success(t *testing.T) {
	// Set environment variables for the test
	err := os.Setenv("DATABASE_URL", "postgres://localhost:5432/testdb")
	require.NoError(t, err)

	err = os.Setenv("KAVENEGAR_API_KEY", "test-api-key")
	require.NoError(t, err)

	err = os.Setenv("KAVENEGAR_TEMPLATE", "test-template")
	require.NoError(t, err)

	// Test loading the config
	cfg, err := config.Load()
	require.NoError(t, err)

	// Assert the correct values were loaded from .env
	assert.Equal(t, "postgres://localhost:5432/testdb", cfg.DatabaseURL)
	assert.Equal(t, "test-api-key", cfg.KavenegarAPIKey)
	assert.Equal(t, "test-template", cfg.KavenegarTemplate)
}

func TestLoadConfig_MissingVariables(t *testing.T) {
	// ایجاد یک فایل .env با متغیرهای ناقص
	err := os.Setenv("DATABASE_URL", "postgres://localhost:5432/testdb")
	require.NoError(t, err)

	err = os.Setenv("KAVENEGAR_API_KEY", "")
	require.NoError(t, err)

	err = os.Setenv("KAVENEGAR_TEMPLATE", "")
	require.NoError(t, err)

	// تلاش برای بارگذاری پیکربندی
	_, err = config.Load()

	// بررسی اینکه خطا به درستی برگشت داده شده است
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "تنظیمات ضروری از فایل .env")
}

func TestLoadConfig_InvalidEnv(t *testing.T) {
	// تنظیم متغیرهای محیطی نادرست برای تست
	err := os.Setenv("DATABASE_URL", "invalid-url")
	require.NoError(t, err)

	// تلاش برای بارگذاری پیکربندی با مقادیر نادرست
	_, err = config.Load()

	// بررسی اینکه خطای مربوط به متغیرهای نادرست برگشت داده شده است
	assert.Error(t, err)
	
	assert.Contains(t, err.Error(), "تنظیمات ضروری از فایل .env")
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	// حذف فایل .env برای تست
	_, err := os.Stat(".env")
	if err == nil {
		err := os.Remove(".env")  // مطمئن می‌شویم که فایل .env حذف شده باشد
		require.NoError(t, err)
	}

	// تلاش برای بارگذاری پیکربندی
	_, err = config.Load()

	// بررسی اینکه خطا باید برگشت داده شود
	assert.Error(t, err)
	fmt.Println(err.Error())
	assert.Contains(t, err.Error(), "تنظیمات ضروری از فایل .env")
}


