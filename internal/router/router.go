package router

import (
	"net/http"
	"database/sql"
	userHandler "github.com/rezalaal/coral/internal/user/handler"
	userRepoInterfaces "github.com/rezalaal/coral/internal/user/repository/interfaces"
	authHandler "github.com/rezalaal/coral/internal/auth/handler"
	authRepoInterfaces "github.com/rezalaal/coral/internal/auth/repository/interfaces"
	"github.com/rezalaal/coral/internal/auth/services"
	"github.com/rezalaal/coral/config"
)

func NewRouter(db *sql.DB, userRepo userRepoInterfaces.UserRepository, otpRepo authRepoInterfaces.OTPRepository) http.Handler {
	mux := http.NewServeMux()

	// Handlers
	userHandler := userHandler.NewUserHandler(userRepo)

	// ایجاد سرویس Kavenegar
	cfg, err := config.Load()
	if err != nil {
		panic("خطا در خواندن تنظیمات .env") // یا می‌توانید یک خطای مناسب مدیریت کنید
	}
	kavenegarService := services.NewKavenegarService(cfg.KavenegarAPIKey)

	// ساخت OTPService
	otpService := services.NewOTPService(otpRepo, kavenegarService) // ارسال KavenegarService به OTPService

	// ایجاد OTPHandler
	otpHandler := authHandler.NewOTPHandler(otpService)

	// روت‌ها
	mux.HandleFunc("/users", userHandler.GetUsers)           // GET
	mux.HandleFunc("/users/create", userHandler.CreateUser)  // POST

	mux.HandleFunc("/otp/send", otpHandler.SendOTP)        // POST
	mux.HandleFunc("/otp/verify", otpHandler.VerifyOTP)    // POST

	return mux
}
