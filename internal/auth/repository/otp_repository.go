package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type OTPRepository struct {
	DB *sqlx.DB
}

func NewOTPRepository(db *sqlx.DB) *OTPRepository {
	return &OTPRepository{DB: db}
}

// ذخیره OTP در دیتابیس
func (r *OTPRepository) SaveOTP(mobile, otp string) error {
	_, err := r.DB.Exec("INSERT INTO otp (mobile, otp, created_at) VALUES ($1, $2, $3)", mobile, otp, time.Now())
	if err != nil {
		return fmt.Errorf("خطا در ذخیره OTP: %v", err)
	}
	return nil
}

// دریافت OTP از دیتابیس
func (r *OTPRepository) GetOTP(mobile string) (string, error) {
	var otp string
	err := r.DB.Get(&otp, "SELECT otp FROM otp WHERE mobile=$1 ORDER BY created_at DESC LIMIT 1", mobile)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("OTP برای این شماره موجود نیست")
		}
		return "", fmt.Errorf("خطا در دریافت OTP: %v", err)
	}
	return otp, nil
}
