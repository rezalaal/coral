package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(c *Container) http.Handler {
	r := mux.NewRouter()

	// OTP Routes
	r.HandleFunc("/otp/send", c.OTPHandler.Send).Methods("POST")
	r.HandleFunc("/otp/verify", c.OTPHandler.Verify).Methods("POST")

	// User Routes
	r.HandleFunc("/register", c.UserHandler.RegisterUser)

	r.HandleFunc("/user/register", c.UserHandler.Register).Methods("POST")
	r.HandleFunc("/user/login", c.UserHandler.Login).Methods("POST")
	r.Handle("/user/me", c.JWTMiddleware.Verify(http.HandlerFunc(c.UserHandler.Me))).Methods("GET")
	r.HandleFunc("/auth/login-with-otp", c.AuthHandler.LoginWithOTP).Methods("POST")

	return r
}
