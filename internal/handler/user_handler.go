// internal/handler/user_handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/rezalaal/coral/internal/models"
	"github.com/rezalaal/coral/internal/repository/interfaces"
)

type UserHandler struct {
	Repo interfaces.UserRepository
}

func NewUserHandler(repo interfaces.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

// دریافت لیست کاربران
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.List() // بدون ctx
	if err != nil {
		http.Error(w, "خطا در دریافت کاربران", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "خطای تجزیه‌ی JSON", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Mobile) == "" {
		http.Error(w, "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند", http.StatusBadRequest)
		return
	}

	err := h.Repo.Create(&user)
	if err != nil {
		http.Error(w, "خطا در ایجاد کاربر", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
