package services

import (
	"fmt"
	"github.com/kavenegar/kavenegar-go"
)

type VerifyLookupParam struct {
	Token   string
	Token2  string
	Token3  string
	Tokens  map[string]string
	Type    string
}

// KavenegarService ساختاری برای سرویس Kavenegar
type KavenegarService struct {
	client *kavenegar.Client
	verify *kavenegar.VerifyService
}

// NewKavenegarService برای ایجاد یک نمونه از سرویس Kavenegar
func NewKavenegarService(apiKey string) *KavenegarService {
	client := kavenegar.NewClient(apiKey)
	verify := kavenegar.NewVerifyService(client)
	return &KavenegarService{client: client, verify: verify}
}

// VerifyLookup متد VerifyLookup برای ارسال OTP با استفاده از Kavenegar
func (s *KavenegarService) VerifyLookup(mobile, token, template string, params *kavenegar.VerifyLookupParam) (*VerifyResponse, error) {
	// استفاده از متد Verify.Lookup از VerifyService
	kavenegarResponse, err := s.verify.Lookup(mobile, token, template, params)
	if err != nil {
		return nil, fmt.Errorf("خطا در ارسال OTP با استفاده از متد lookup: %v", err)
	}

	// ساختار جدید باید با کدهای جدید تطابق داشته باشد
	response := &VerifyResponse{
		MessageID:  int64(kavenegarResponse.MessageID), // تبدیل به int64
		Message:    kavenegarResponse.Message,
		Status:     int(kavenegarResponse.Status), // تبدیل به int
		StatusText: kavenegarResponse.StatusText,
		Sender:     kavenegarResponse.Sender,
		Receptor:   kavenegarResponse.Receptor,
		Date:       int64(kavenegarResponse.Date),
		Cost:       kavenegarResponse.Cost,
	}

	return response, nil
}
