package handler

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	userusecase "coral/internal/application/usecase/user"
	"coral/internal/interfaces/http/middleware"
	"coral/internal/utils"
)

type UserHandler struct {
	RegisterUC *userusecase.RegisterUseCase
	LoginUC    *userusecase.LoginUseCase
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	u, err := h.RegisterUC.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.LoginUC.Execute(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	if userID == nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"status":  "authenticated",
	})
}


func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// اعتبارسنجی اولیه
	if name == "" || email == "" || password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `<div class="text-red-600 text-sm">تمام فیلدها الزامی هستند.</div>`)
		return
	}

	if !utils.IsValidEmail(email) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, `<div class="text-red-600 text-sm">ایمیل معتبر نیست.</div>`)
		return
	}

	// بررسی تکراری نبودن ایمیل و ذخیره در دیتابیس
	err := h.UserService.CreateUser(name, email, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `<div class="text-red-600 text-sm">%s</div>`, html.EscapeString(err.Error()))
		return
	}

	// موفقیت
	fmt.Fprint(w, `<div class="text-green-600 text-sm">ثبت‌نام با موفقیت انجام شد.</div>`)
}
