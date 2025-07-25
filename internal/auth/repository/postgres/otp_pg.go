package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/rezalaal/coral/internal/auth/repository/interfaces"
)

// OTPRepository ساختار برای تعامل با دیتابیس برای OTP
type OTPRepository struct {
	DB *sql.DB
}

// NewOTPRepository ایجاد یک نمونه جدید از OTPRepository
func NewOTPRepository(db *sql.DB) *OTPRepository {
	return &OTPRepository{DB: db}
}

// ذخیره OTP در دیتابیس
func (r *OTPRepository) SaveOTP(mobile, otp string) error {
	query := "INSERT INTO otp_codes (mobile, code, created_at, expires_at) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, mobile, otp, time.Now(), time.Now().Add(10*time.Minute))
	if err != nil {
		return fmt.Errorf("خطا در ذخیره OTP: %v", err)
	}
	return nil
}

// دریافت OTP از دیتابیس
func (r *OTPRepository) GetOTP(mobile string) (string, error) {
	var otp string
	query := "SELECT code FROM otp_codes WHERE mobile = $1 ORDER BY created_at DESC LIMIT 1"
	err := r.DB.QueryRow(query, mobile).Scan(&otp)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("OTP برای این شماره موجود نیست")
		}
		return "", fmt.Errorf("خطا در دریافت OTP: %v", err)
	}
	return otp, nil
}
