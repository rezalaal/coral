package services

import (
	"github.com/kavenegar/kavenegar-go"
	"fmt"
)

// KavenegarService ساختاری است که سرویس Kavenegar را پیاده‌سازی می‌کند
type KavenegarService struct {
	client *kavenegar.Client
}

// NewKavenegarService برای ایجاد یک نمونه از سرویس Kavenegar
func NewKavenegarService(apiKey string) *KavenegarService {
	client := kavenegar.NewClient(apiKey)
	return &KavenegarService{client: client}
}

// VerifyLookup متد VerifyLookup برای ارسال OTP با استفاده از Kavenegar
func (s *KavenegarService) VerifyLookup(mobile, token, template string, params *kavenegar.VerifyLookupParam) (*kavenegar.APIResponse, error) {
	return s.client.Verify.Lookup(mobile, token, template, params)
}
