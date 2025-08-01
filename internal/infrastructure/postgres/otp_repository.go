package postgres

import (
	"database/sql"
	"errors"
)

type OTPRepo struct {
	DB *sql.DB
}

func NewOTPRepository(db *sql.DB) *OTPRepo {
	return &OTPRepo{DB: db}
}

func (r *OTPRepo) Save(mobile, code string, expiresAt int64) error {
	_, err := r.DB.Exec("INSERT INTO otps (mobile, code, expires_at) VALUES ($1, $2, $3) ON CONFLICT (mobile) DO UPDATE SET code=$2, expires_at=$3", mobile, code, expiresAt)
	return err
}

func (r *OTPRepo) Get(mobile string) (string, int64, error) {
	var code string
	var expires int64
	err := r.DB.QueryRow("SELECT code, expires_at FROM otps WHERE mobile = $1", mobile).Scan(&code, &expires)
	if err != nil {
		return "", 0, errors.New("not found")
	}
	return code, expires, nil
}
