package models

// مدل درخواست OTP
type OTPRequest struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}
