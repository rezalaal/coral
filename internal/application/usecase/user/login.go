package user

import (
	"errors"
	"github.com/rezalaal/coral/internal/domain/user"
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

	if !user.ComparePassword(u.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return uc.Token.Generate(u.ID)
}
