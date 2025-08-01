package postgres

import (
	"database/sql"
	"errors"
	"github.com/rezalaal/coral/internal/domain/user"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(u *user.User) error {
	return r.DB.QueryRow(
		"INSERT INTO users (mobile) VALUES ($1) RETURNING id",
		u.Mobile,
	).Scan(&u.ID)
}


func (r *UserRepo) FindByEmail(email string) (*user.User, error) {
	var u user.User
	err := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) FindByMobile(mobile string) (*user.User, error) {
	var u user.User
	err := r.DB.QueryRow("SELECT id, name, email, mobile FROM users WHERE mobile = $1", mobile).
		Scan(&u.ID, &u.Name, &u.Email, &u.Mobile)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}
