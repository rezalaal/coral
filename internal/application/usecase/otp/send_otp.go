package otp

import (
	"time"
	"github.com/rezalaal/coral/internal/domain/otp"
)

type SendOTPUseCase struct {
	Sender   otp.Sender
	Repo     otp.Repository
	Generate func() string
	Throttle otp.ThrottleRepository // اضافه‌شده برای محدودسازی
}

const MaxRequestsPerDay = 5

func (uc *SendOTPUseCase) Execute(mobile string) error {
	// بررسی محدودیت
	count, err := uc.Throttle.CountRecentRequests(mobile, 24*time.Hour)
	if err != nil {
		return err
	}
	if count >= MaxRequestsPerDay {
		return errors.New("OTP request limit reached for today")
	}

	// ادامه فرآیند ارسال
	code := uc.Generate()
	if err := uc.Sender.SendOTP(mobile, code); err != nil {
		return err
	}

	expires := time.Now().Add(2 * time.Minute).Unix()
	if err := uc.Repo.Save(mobile, code, expires); err != nil {
		return err
	}

	return uc.Throttle.LogRequest(mobile) // ثبت درخواست موفق
}

