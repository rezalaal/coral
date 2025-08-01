package user

import (
	"errors"

	"coral/internal/models"
	"coral/internal/repository/interfaces"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo interfaces.UserRepository
}

func NewService(repo interfaces.UserRepository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) CreateUser(name, email, password string) error {
	// بررسی ایمیل تکراری
	existing, err := s.Repo.FindByEmail(email)
	if err != nil && !errors.Is(err, interfaces.ErrUserNotFound) {
		return err
	}
	if existing != nil {
		return errors.New("این ایمیل قبلاً ثبت شده است")
	}

	// هش کردن رمز عبور
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	return s.Repo.Create(user)
}
