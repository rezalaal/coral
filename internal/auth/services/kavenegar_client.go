package services

import "github.com/kavenegar/kavenegar-go"

// تعریف interface برای سرویس Kavenegar
type KavenegarClient interface {
	VerifyLookup(mobile, token, template string, params *kavenegar.VerifyLookupParam) (*kavenegar.APIResponse, error)
}
