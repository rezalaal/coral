package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/rezalaal/coral/internal/user/models"
	"github.com/rezalaal/coral/internal/user/repository/interfaces"
	"github.com/rezalaal/coral/internal/utils"
	"github.com/rezalaal/coral/internal/validation"
)

type UserHandler struct {
	Repo interfaces.UserRepository
}

func NewUserHandler(repo interfaces.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "خطای تجزیه‌ی JSON", http.StatusBadRequest)
		return
	}

	// بررسی فیلدهای ضروری
	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Mobile) == "" {
		http.Error(w, "خطا در ایجاد کاربر: فیلدهای ضروری ناقص هستند", http.StatusBadRequest)
		return
	}

	// تبدیل شماره موبایل به انگلیسی در صورت لزوم
	user.Mobile = utils.ConvertPersianToEnglishNumbers(user.Mobile)

	// اعتبارسنجی ورودی‌ها
	if !validation.IsValidName(user.Name) {
		http.Error(w, "نام معتبر نیست", http.StatusBadRequest)
		return
	}

	if len(user.Mobile) != 11 || !validation.IsValidMobile(user.Mobile) {
		http.Error(w, "شماره موبایل نامعتبر است", http.StatusBadRequest)
		return
	}

	// ایجاد کاربر در دیتابیس
	err := h.Repo.Create(&user)
	if err != nil {
		log.Printf("Error creating user: %v", err)  // ثبت خطا در لاگ
		http.Error(w, "خطا در ایجاد کاربر", http.StatusInternalServerError)
		return
	}

	// ارسال پاسخ موفقیت‌آمیز
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Repo.List()
	if err != nil {
		http.Error(w, "خطا در دریافت کاربران", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
