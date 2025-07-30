// internal/auth/repository/interfaces/otp_repository.go
package interfaces

// OTPRepository interface برای عملیات ذخیره و بازیابی OTP
type OTPRepository interface {
	// ذخیره OTP در دیتابیس
	SaveOTP(mobile string, otp string) error

	// دریافت OTP از دیتابیس
	GetOTP(mobile string) (string, error)
}
