package router

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/rezalaal/coral/config"
	authHandler "github.com/rezalaal/coral/internal/auth/handler"
	authRepoInterfaces "github.com/rezalaal/coral/internal/auth/repository/interfaces"
	"github.com/rezalaal/coral/internal/auth/services"
	userHandler "github.com/rezalaal/coral/internal/user/handler"
	userRepoInterfaces "github.com/rezalaal/coral/internal/user/repository/interfaces"
)

func NewRouter(db *sql.DB, userRepo userRepoInterfaces.UserRepository, otpRepo authRepoInterfaces.OTPRepository) http.Handler {
	mux := http.NewServeMux()

	// Handlers
	userHandler := userHandler.NewUserHandler(userRepo)

	// خواندن تنظیمات .env
	cfg, err := config.Load()
	if err != nil {
		panic("خطا در خواندن تنظیمات .env")
	}

	// چک کردن متغیر محیطی
	environment := os.Getenv("ENVIRONMENT")
	var otpService *services.OTPService

	if environment == "development" {
		// در حالت development از Mock استفاده می‌کنیم
		otpService = services.NewOTPService(otpRepo, &services.MockKavenegarClient{})
		log.Println("Running in development mode. Using Mock OTP Service.")
	} else {
		// در حالت‌های دیگر از Kavenegar استفاده می‌کنیم
		kavenegarService := services.NewKavenegarService(cfg.KavenegarAPIKey)
		otpService = services.NewOTPService(otpRepo, kavenegarService) // ارسال KavenegarService به OTPService
		log.Println("Running in production mode. Using real Kavenegar Service.")
	}

	// ایجاد OTPHandler
	otpHandler := authHandler.NewOTPHandler(otpService)

	// روت‌ها
	mux.HandleFunc("/users", userHandler.GetUsers)           // GET
	mux.HandleFunc("/users/create", userHandler.CreateUser)  // POST
	
	// مسیر hello برای تست
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request to /hello")
		w.Write([]byte("Hello, World!"))
	})

	mux.HandleFunc("/otp/send", otpHandler.SendOTP)       // POST
	mux.HandleFunc("/otp/verify", otpHandler.VerifyOTP)   // POST

	return mux
}
