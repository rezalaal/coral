package otp

import (
	"time"
	"coral/internal/domain/otp"
)

type VerifyOTPUseCase struct {
	Repo otp.Repository
}

func (uc *VerifyOTPUseCase) Execute(mobile, code string) (bool, error) {
	storedCode, expiresAt, err := uc.Repo.Get(mobile)
	if err != nil {
		return false, err
	}
	if code != storedCode {
		return false, nil
	}
	if time.Now().Unix() > expiresAt {
		return false, nil
	}
	return true, nil
}
