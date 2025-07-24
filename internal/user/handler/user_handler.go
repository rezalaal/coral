package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/rezalaal/coral/internal/user/models"
	"github.com/rezalaal/coral/internal/user/repository/interfaces"
	"github.com/rezalaal/coral/internal/utils"
)

type UserHandler struct {
	Repo interfaces.UserRepository
}

func NewUserHandler(repo interfaces.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func isAlpha(s string) bool {
	// اصلاح regex برای شناسایی حروف فارسی، انگلیسی و فضای خالی
	return regexp.MustCompile(`^[a-zA-Z\x{0600}-\x{06FF}\s]+$`).MatchString(s)
}

func isValidName(name string) bool {
	// بررسی اینکه نام تنها شامل حروف انگلیسی و فارسی باشد
	if !isAlpha(name) {
		return false
	}

	// بررسی طول نام
	if len(name) > 100 || len(name) == 0 {
		return false
	}

	return true
}

func isValidMobile(mobile string) bool {
	// بررسی فرمت شماره موبایل (فقط اعداد و طول دقیق 11 رقم)
	return regexp.MustCompile(`^\d{11}$`).MatchString(mobile)
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
	if !isValidName(user.Name) {
		http.Error(w, "نام معتبر نیست", http.StatusBadRequest)
		return
	}

	if len(user.Mobile) != 11 || !isValidMobile(user.Mobile) {
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
