package handler

import (
	"encoding/json"
	"net/http"
	"github.com/rezalaal/coral/internal/application/usecase/otp"
)

type OTPHandler struct {
	SendUC   *otp.SendOTPUseCase
	VerifyUC *otp.VerifyOTPUseCase
}

func (h *OTPHandler) Send(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Mobile string `json:"mobile"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	err := h.SendUC.Execute(req.Mobile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *OTPHandler) Verify(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Mobile string `json:"mobile"`
		Code   string `json:"code"`
	}
	json.NewDecoder(r.Body).Decode(&req)
	valid, err := h.VerifyUC.Execute(req.Mobile, req.Code)
	if err != nil || !valid {
		http.Error(w, "invalid code", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
