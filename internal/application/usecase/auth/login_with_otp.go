package auth

import (
	"errors"
	"github.com/rezalaal/coral/internal/domain/otp"
	"github.com/rezalaal/coral/internal/domain/user"
	"time"
)

type LoginWithOTPUseCase struct {
	OTPRepo otp.Repository
	UserRepo user.Repository
	TokenGen user.TokenGenerator
}

func (uc *LoginWithOTPUseCase) Execute(mobile string, code string) (string, error) {
	storedCode, expiresAt, err := uc.OTPRepo.Get(mobile)
	if err != nil || storedCode != code || expiresAt < time.Now().Unix() {
		return "", errors.New("invalid or expired code")
	}

	u, err := uc.UserRepo.FindByMobile(mobile)
	if err != nil || u == nil {
		u = &user.User{Mobile: mobile}
		err = uc.UserRepo.Create(u) // ثبت‌نام جدید
		if err != nil {
			return "", err
		}
	}

	return uc.TokenGen.Generate(u.ID)
}
