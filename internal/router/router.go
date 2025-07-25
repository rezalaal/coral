package router

import (
	"net/http"
	"database/sql"
	userHandler "github.com/rezalaal/coral/internal/user/handler"
	userRepoInterfaces "github.com/rezalaal/coral/internal/user/repository/interfaces"
	authHandler "github.com/rezalaal/coral/internal/auth/handler"
	authRepoInterfaces "github.com/rezalaal/coral/internal/auth/repository/interfaces"
)

func NewRouter(db *sql.DB, userRepo userRepoInterfaces.UserRepository, otpRepo authRepoInterfaces.OTPRepository) http.Handler {
	mux := http.NewServeMux()

	// Handlers
	userHandler := userHandler.NewUserHandler(userRepo)
	otpHandler := authHandler.NewOTPHandler(otpRepo) // استفاده از OTP handler مشابه user handler

	// روت‌ها
	mux.HandleFunc("/users", userHandler.GetUsers)           // GET
	mux.HandleFunc("/users/create", userHandler.CreateUser)  // POST

	mux.HandleFunc("/otp/send", otpHandler.SendOTP)        // POST
	mux.HandleFunc("/otp/verify", otpHandler.VerifyOTP)    // POST

	return mux
}
