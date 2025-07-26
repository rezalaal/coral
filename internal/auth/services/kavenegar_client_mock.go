// internal/auth/services/kavenegar_client_mock.go

package services

// MockKavenegarClient برای شبیه‌سازی رفتار KavenegarClient در تست‌ها
type MockKavenegarClient struct{}

// VerifyLookup شبیه‌سازی متد VerifyLookup
func (m *MockKavenegarClient) VerifyLookup(mobile, token, template string, params *VerifyLookupParam) (*VerifyResponse, error) {
	// شبیه‌سازی پاسخ موفق
	return &VerifyResponse{
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
