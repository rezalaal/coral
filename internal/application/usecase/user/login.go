package user

import (
	"errors"
	"coral/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase struct {
	Repo  user.Repository
	Token user.TokenGenerator
}

func (uc *LoginUseCase) Execute(email, password string) (string, error) {
	u, err := uc.Repo.FindByEmail(email)
	if err != nil || u == nil {
		return "", errors.New("invalid credentials")
	}

	if !ComparePassword(u.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return uc.Token.Generate(u.ID)
}

func ComparePassword(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}