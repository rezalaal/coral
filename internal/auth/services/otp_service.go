package service

import (
	"fmt"
	"os"

	"github.com/kavenegar/kavenegar-go"
	"github.com/rezalaal/coral/config"
	"github.com/rezalaal/coral/internal/auth/repository"
)

type OTPService struct {
	Repository *repository.OTPRepository
}

func NewOTPService(repository *repository.OTPRepository) *OTPService {
	return &OTPService{Repository: repository}
}

// ارسال OTP با استفاده از متد lookup
func (s *OTPService) SendOTP(mobile string) error {
	cfg ,err := config.Load()
	if err != nil {
		return fmt.Errorf("خطا در خواندن تنظیمات  .env: %v", err)
	}
	apiKey := cfg.KavenegarAPIKey
	client, err := kavenegar.NewClient(apiKey)
	if err != nil {
		return fmt.Errorf("خطا در ایجاد مشتری Kavenegar: %v", err)
	}

	// دریافت کد تصادفی برای OTP
	token := "852596" // این باید به طور داینامیک ایجاد شود (مثلاً یک کد تصادفی)

	// ارسال OTP به شماره موبایل با استفاده از lookup
	_, err = client.Verify.Lookup(mobile, token, cfg.KavenegarTemplate)
	if err != nil {
		return fmt.Errorf("خطا در ارسال OTP با استفاده از متد lookup: %v", err)
	}

	// ذخیره کد OTP در دیتابیس
	err = s.Repository.SaveOTP(mobile, token) // ذخیره کد OTP
	if err != nil {
		return fmt.Errorf("خطا در ذخیره OTP: %v", err)
	}

	return nil
}

// تایید OTP
func (s *OTPService) VerifyOTP(mobile, code string) (bool, error) {
	// بررسی صحت کد OTP از دیتابیس
	storedCode, err := s.Repository.GetOTP(mobile)
	if err != nil {
		return false, fmt.Errorf("خطا در دریافت OTP: %v", err)
	}

	return storedCode == code, nil
}
