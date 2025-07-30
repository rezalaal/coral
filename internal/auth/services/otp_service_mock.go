package services

import "fmt"

// MockOTPService شبیه‌ساز برای سرویس ارسال OTP در محیط توسعه
type MockOTPService struct{}

// SendOTP شبیه‌سازی ارسال OTP
func (m *MockOTPService) SendOTP(mobile string) error {
	// اینجا می‌توانیم کد را برای تست ارسال کنیم
	fmt.Println("OTP sent (Mock) to mobile:", mobile)
	return nil
}

// VerifyOTP شبیه‌سازی تایید OTP
func (m *MockOTPService) VerifyOTP(mobile, code string) (bool, error) {
	// اینجا می‌توانیم یک کد پیش‌فرض برای تایید استفاده کنیم
	return code == "852596", nil
}
