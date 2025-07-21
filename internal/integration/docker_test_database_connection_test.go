package integration_test

import (
	"database/sql"
	"os"
	
	"path/filepath"

	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	// تنظیمات اتصال به پایگاه داده
	rootPath, _ := filepath.Abs(filepath.Join("..", "..", ".env"))
	err := godotenv.Load(rootPath)
	assert.NoError(t, err)
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	defer db.Close()

	// تست اتصال به پایگاه داده
	err = db.Ping()
	require.NoError(t, err, "Failed to connect to database")

	// اگر اتصال موفقیت‌آمیز بود، باید پیامی مشابه این برگردانده شود
	t.Log("Successfully connected to the database")
}
