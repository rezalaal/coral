package integration

import (
	"database/sql"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	"github.com/rezalaal/coral/internal/db"
	userPG "github.com/rezalaal/coral/internal/user/repository/postgres"
	otpPG  "github.com/rezalaal/coral/internal/auth/repository/postgres"
	"github.com/rezalaal/coral/internal/router"
)

// اتصال به دیتابیس با بررسی خطا
func connectTestDB(t *testing.T) *sql.DB {
	rootPath, _ := filepath.Abs(filepath.Join("..", "..", ".env"))
	err := godotenv.Load(rootPath)
	assert.NoError(t, err)
	
	conn, err := db.Connect()
	assert.NoError(t, err)
	return conn
}

// پاک‌سازی دیتابیس (در حال حاضر فقط جدول users)
func cleanupDB(t *testing.T, db *sql.DB) {
	_, err := db.Exec("DELETE FROM users")
	assert.NoError(t, err)
}

func SetupTestServer(t *testing.T) (*httptest.Server, func()) {
	dbConn := connectTestDB(t)
	cleanupDB(t, dbConn)

	userRepo := userPG.NewUserPG(dbConn)
	// ایجاد OTPRepository برای تست
	otpRepo := otpPG.NewOTPRepository(dbConn)

	// ارسال userRepo و otpRepo به NewRouter
	r := router.NewRouter(dbConn, userRepo, otpRepo)
	server := httptest.NewServer(r)

	teardown := func() {
		dbConn.Close()
		server.Close()
	}

	return server, teardown
}