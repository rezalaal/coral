package kavenegar

import (
	"fmt"
)

type KavenegarSender struct {
	APIKey string
}

func (s *KavenegarSender) SendOTP(mobile, code string) error {
	fmt.Printf("Sending OTP %s to %s via Kavenegar\n", code, mobile)
	// کد واقعی ارتباط با API Kavenegar اینجا قرار می‌گیرد
	return nil
}
