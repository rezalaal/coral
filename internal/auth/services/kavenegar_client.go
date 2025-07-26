package services

import "github.com/kavenegar/kavenegar-go"

// تعریف struct برای پاسخ VerifyLookup
type VerifyResponse struct {
	MessageID   int64  `json:"messageid"`
	Message     string `json:"message"`
	Status      int    `json:"status"`
	StatusText  string `json:"statustext"`
	Sender      string `json:"sender"`
	Receptor    string `json:"receptor"`
	Date        int64  `json:"date"`
	Cost        int    `json:"cost"`
}


// تعریف interface برای سرویس Kavenegar
type KavenegarClient interface {
	VerifyLookup(mobile, token, template string, params *kavenegar.VerifyLookupParam) (*VerifyResponse, error)
}
