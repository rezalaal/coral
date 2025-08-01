package user

import (
	"errors"
	"coral/internal/domain/user"
)

type RegisterUseCase struct {
	Repo user.Repository
}

func (uc *RegisterUseCase) Execute(name, email, password string) (*user.User, error) {
	existing, _ := uc.Repo.FindByEmail(email)
	if existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed := HashPassword(password)
	u := &user.User{
		Name:     name,
		Email:    email,
		Password: hashed,
	}
	err := uc.Repo.Create(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

