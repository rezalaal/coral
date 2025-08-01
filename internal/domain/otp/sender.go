package otp

type Sender interface {
	SendOTP(mobile, code string) error
}
