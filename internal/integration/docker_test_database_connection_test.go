// internal/integration/docker_test_database_connection.go
package integration_test

import (
	"database/sql"
	"testing"
	
	_ "github.com/lib/pq" // PostgreSQL driver
	"coral/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatabaseConnection(t *testing.T) {
	// بارگذاری تنظیمات از فایل config
	cfg, err := config.Load()
	assert.NoError(t, err)

	// استفاده از DATABASE_URL از config
	connStr := cfg.DatabaseURL
	db, err := sql.Open("postgres", connStr)
	require.NoError(t, err)
	defer db.Close()

	// تست اتصال به پایگاه داده
	err = db.Ping()
	require.NoError(t, err, "Failed to connect to database")

	// اگر اتصال موفقیت‌آمیز بود، باید پیامی مشابه این برگردانده شود
	t.Log("Successfully connected to the database")
}
