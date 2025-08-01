package usecase_test

import (
	"testing"
	"time"

	"github.com/rezalaal/coral/internal/application/usecase/otp"
	"github.com/rezalaal/coral/internal/tests/mocks"
)

type fakeSender struct{}

func (s *fakeSender) SendOTP(mobile, code string) error {
	return nil
}

func TestSendOTP(t *testing.T) {
	repo := mocks.NewInMemoryOTPRepo()
	sender := &fakeSender{}

	uc := otp.SendOTPUseCase{
		Sender:   sender,
		Repo:     repo,
		Generate: func() string { return "123456" },
	}

	err := uc.Execute("09120000000")
	if err != nil {
		t.Fatal("unexpected error:", err)
	}

	code, exp, err := repo.Get("09120000000")
	if code != "123456" || exp < time.Now().Unix() {
		t.Error("invalid saved otp")
	}
}
