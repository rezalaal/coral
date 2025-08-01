package router

import (
	"database/sql"
	"github.com/rezalaal/coral/internal/interfaces/http/handler"
	"github.com/rezalaal/coral/internal/interfaces/http/middleware"
	"github.com/rezalaal/coral/internal/infrastructure/kavenegar"
	"github.com/rezalaal/coral/internal/infrastructure/jwt"
	"github.com/rezalaal/coral/internal/infrastructure/postgres"
	otpusecase "github.com/rezalaal/coral/internal/application/usecase/otp"
	userusecase "github.com/rezalaal/coral/internal/application/usecase/user"
	"pkg/utils"
)

type Container struct {
	DB            *sql.DB
	OTPHandler    *handler.OTPHandler
	UserHandler   *handler.UserHandler
	JWTMiddleware *middleware.JWTMiddleware
}

func NewContainer(db *sql.DB) *Container {

	jwtSecret := os.Getenv("JWT_SECRET")
	kavenegarKey := os.Getenv("KAVENEGAR_API_KEY")

	// OTP
	otpRepo := postgres.NewOTPRepository(db)
	otpSender := &kavenegar.KavenegarSender{APIKey: kavenegarKey}
	throttleRepo := postgres.NewOTPThrottleRepo(db)
	sendOTP := &otpusecase.SendOTPUseCase{
		Sender:   otpSender,
		Repo:     otpRepo,
		Generate: utils.GenerateOTPCode,
		Throttle: throttleRepo,
	}

	verifyOTP := &otpusecase.VerifyOTPUseCase{Repo: otpRepo}
	otpHandler := &handler.OTPHandler{
		SendUC:   sendOTP,
		VerifyUC: verifyOTP,
	}

	// User
	userRepo := postgres.NewUserRepository(db)
	registerUC := &userusecase.RegisterUseCase{Repo: userRepo}
	tokenGen := &jwt.JWTToken{Secret: jwtSecret}
	loginUC := &userusecase.LoginUseCase{
		Repo:  userRepo,
		Token: tokenGen,
	}
	userHandler := &handler.UserHandler{
		RegisterUC: registerUC,
		LoginUC:    loginUC,
	}

	loginWithOTP := &auth.LoginWithOTPUseCase{
		OTPRepo:   otpRepo,
		UserRepo:  userRepo,
		TokenGen:  tokenGen,
	}
	authHandler := &handler.AuthHandler{
		LoginWithOTP: loginWithOTP,
	}
	c.AuthHandler = authHandler


	// JWT Middleware	
	jwtMiddleware := &middleware.JWTMiddleware{Secret: jwtSecret}
	return &Container{
		DB:            db,
		OTPHandler:    otpHandler,
		UserHandler:   userHandler,
		JWTMiddleware: jwtMiddleware,
	}
}
