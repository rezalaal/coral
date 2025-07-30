// internal/auth/services/kavenegar_client_mock.go
package services

import (
	"github.com/rezalaal/coral/internal/auth/services"
	"github.com/kavenegar/kavenegar-go"
	"log"
)

// MockKavenegarClient برای شبیه‌سازی رفتار KavenegarClient در تست‌ها
type MockKavenegarClient struct{}

// VerifyLookup شبیه‌سازی متد VerifyLookup
func (m *MockKavenegarClient) VerifyLookup(mobile, token, template string, params *kavenegar.VerifyLookupParam) (*services.VerifyResponse, error) {
	// شبیه‌سازی پاسخ موفق
	log.Printf("MockKavenegarClient: VerifyLookup called with mobile: %s, token: %s, template: %s, params: %+v", mobile, token, template, params)
	return &services.VerifyResponse{
		MessageID: 123456,
		Message:   "کد تایید عضویت: 852596",
		Status:    200,
		StatusText: "ارسال به مخابرات",
		Sender:    "10004346",
		Receptor:  mobile,
		Date:      1356619709,
		Cost:      120,
	}, nil
}
