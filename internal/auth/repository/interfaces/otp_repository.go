package interfaces

// OTPRepository interface برای عملیات ذخیره و بازیابی OTP
type OTPRepository interface {
	SaveOTP(mobile, otp string) error
	GetOTP(mobile string) (string, error)
}
