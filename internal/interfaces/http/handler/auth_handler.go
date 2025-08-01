type AuthHandler struct {
	LoginWithOTP *auth.LoginWithOTPUseCase
}

func (h *AuthHandler) LoginWithOTP(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Mobile string `json:"mobile"`
		Code   string `json:"code"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.LoginWithOTP.Execute(req.Mobile, req.Code)
	if err != nil {
		http.Error(w, "invalid or expired OTP", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
