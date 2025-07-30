// internal/auth/services/otp_service.go
package services

import (
	"fmt"
	"log"
	"github.com/rezalaal/coral/internal/auth/repository/interfaces"
	"github.com/kavenegar/kavenegar-go"
	"github.com/rezalaal/coral/config"
)

type OTPService struct {
	Repository      interfaces.OTPRepository
	KavenegarClient KavenegarClient // وابستگی به KavenegarClient
}

// NewOTPService برای ایجاد یک نمونه از سرویس OTP
func NewOTPService(repository interfaces.OTPRepository, kavenegarClient KavenegarClient) *OTPService {
	log.Println("Created new OTP Service with Repository:", repository, "and KavenegarClient:", kavenegarClient)
	return &OTPService{Repository: repository, KavenegarClient: kavenegarClient}
}

// SendOTP برای ارسال OTP به شماره موبایل
func (s *OTPService) SendOTP(mobile string) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("خطا در خواندن تنظیمات .env: %v", err)
	}

	// دریافت کد تصادفی برای OTP
	token := "852596" // این باید به طور داینامیک ایجاد شود (مثلاً یک کد تصادفی)

	// ارسال OTP به شماره موبایل با استفاده از verify lookup
	params := &kavenegar.VerifyLookupParam{
		Tokens: map[string]string{
			"token": token,
		},
		Type: kavenegar.Type_VerifyLookup_Sms, // تایپ پیامک
	}

	// استفاده از KavenegarClient برای ارسال OTP
	verifyResponse, err := s.KavenegarClient.VerifyLookup(mobile, token, cfg.KavenegarTemplate, params)
	if err != nil {
		return fmt.Errorf("خطا در ارسال OTP با استفاده از متد Verify.Lookup: %v", err)
	}

	// بررسی وضعیت ارسال OTP از Kavenegar
	if verifyResponse.Status != 200 {
		return fmt.Errorf("خطا در ارسال OTP: وضعیت ارسال پیامک: %d", verifyResponse.Status)
	}

	// ذخیره کد OTP در دیتابیس
	err = s.Repository.SaveOTP(mobile, token) // ذخیره کد OTP
	if err != nil {
		return fmt.Errorf("خطا در ذخیره OTP: %v", err)
	}

	return nil
}

// VerifyOTP برای تایید OTP
func (s *OTPService) VerifyOTP(mobile, code string) (bool, error) {
	storedCode, err := s.Repository.GetOTP(mobile)
	if err != nil {
		return false, fmt.Errorf("خطا در دریافت OTP: %v", err)
	}

	return storedCode == code, nil
}
