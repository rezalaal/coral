package postgres

import (
	"database/sql"
	"coral/internal/domain/otp"
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
	err := r.DB.QueryRow("SELECT code, expires FROM otps WHERE mobile = $1", mobile).
		Scan(&code, &expires)
	if err != nil {
		return "", 0, err
	}
	return code, expires, nil
}

func (r *OTPRepo) Find(mobile string) (*otp.OTP, error) {
	var o otp.OTP
	err := r.DB.QueryRow("SELECT mobile, code, expires FROM otps WHERE mobile = $1", mobile).
		Scan(&o.Mobile, &o.Code, &o.Expires)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &o, nil
}
