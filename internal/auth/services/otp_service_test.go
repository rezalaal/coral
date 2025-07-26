// internal/auth/services/otp_service_test.go

package services

import (
	"net/http"
	"testing"

	"github.com/rezalaal/coral/internal/auth/repository/postgres"
	"github.com/rezalaal/coral/internal/integration"
	"github.com/stretchr/testify/assert"
	authService "github.com/rezalaal/coral/internal/auth/services"
)

func TestSendOTP(t *testing.T) {
	// راه‌اندازی سرور تست و دیتابیس واقعی
	server, dbConn, teardown := integration.SetupTestServer(t)
	defer teardown()

	// تعریف شماره موبایل برای ارسال OTP
	mobile := "09120000001" // شماره موبایل واقعی برای تست

	// استفاده از MockKavenegarClient
	mockKavenegarClient := &authService.MockKavenegarClient{}
	otpRepo := postgres.NewOTPRepository(dbConn)
	otpService := authService.NewOTPService(otpRepo, mockKavenegarClient)

	// ارسال OTP
	err := otpService.SendOTP(mobile)

	// بررسی اینکه ارسال OTP بدون خطا انجام شد
	assert.NoError(t, err)

	// اجرای درخواست برای بررسی
	resp, err := http.Get(server.URL + "/otp/send?mobile=" + mobile)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// پاکسازی داده‌ها بعد از تست
	integration.CleanupDB(t, dbConn)
}

func TestVerifyOTP(t *testing.T) {
	// راه‌اندازی سرور تست و دیتابیس واقعی
	server, dbConn, teardown := integration.SetupTestServer(t)
	defer teardown()

	// ارسال OTP به شماره موبایل
	mobile := "09120000001" // شماره موبایل واقعی برای تست
	otpService := authService.NewOTPService(postgres.NewOTPRepository(dbConn), &authService.MockKavenegarClient{})

	// ارسال OTP به موبایل
	err := otpService.SendOTP(mobile)
	assert.NoError(t, err)

	// تایید OTP
	code := "852596" // کد OTP ارسال‌شده
	valid, err := otpService.VerifyOTP(mobile, code)
	assert.NoError(t, err)
	assert.True(t, valid)

	// پاکسازی داده‌ها بعد از تست
	integration.CleanupDB(t, dbConn)
}
