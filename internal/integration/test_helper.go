package integration

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"coral/internal/db"
	"coral/internal/router"
)

// اتصال به دیتابیس با بررسی خطا
func connectTestDB(t *testing.T) *sql.DB {
	conn, err := db.Connect()
	assert.NoError(t, err)
	return conn
}

// پاک‌سازی دیتابیس (در حال حاضر فقط جدول users)
func CleanupDB(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	assert.NoError(t, err)
}

func SetupTestServer(t *testing.T) (*httptest.Server, *sql.DB, func()) {
	dbConn := connectTestDB(t)
	CleanupDB(t, dbConn)

	// ✅ ساخت container
	container := router.NewContainer(dbConn)

	// ✅ استفاده از container
	r := router.NewRouter(container)
	server := httptest.NewServer(r)

	teardown := func() {
		dbConn.Close()
		server.Close()
	}

	return server, dbConn, teardown
}

