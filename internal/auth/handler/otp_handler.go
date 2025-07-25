package handler

import (
	"encoding/json"
	"net/http"
	"github.com/rezalaal/coral/internal/auth/services"
)

type OTPHandler struct {
	OTPService *services.OTPService
}

func NewOTPHandler(otpService *services.OTPService) *OTPHandler {
	return &OTPHandler{OTPService: otpService}
}

// درخواست ارسال OTP
func (h *OTPHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "خطای تجزیه‌ی JSON", http.StatusBadRequest)
		return
	}

	mobile := request["mobile"]

	// ارسال OTP
	err := h.OTPService.SendOTP(mobile)
	if err != nil {
		http.Error(w, "خطا در ارسال OTP", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OTP ارسال شد"})
}

// تایید OTP
func (h *OTPHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {
	var request map[string]string
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "خطای تجزیه‌ی JSON", http.StatusBadRequest)
		return
	}

	mobile := request["mobile"]
	code := request["code"]

	// تایید OTP
	valid, err := h.OTPService.VerifyOTP(mobile, code)
	if err != nil || !valid {
		http.Error(w, "کد تایید نامعتبر است", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "OTP تایید شد"})
}
