package postgres

import (
	"database/sql"
	"time"
)

type OTPThrottleRepo struct {
	DB *sql.DB
}

func NewOTPThrottleRepo(db *sql.DB) *OTPThrottleRepo {
	return &OTPThrottleRepo{DB: db}
}

func (r *OTPThrottleRepo) CountRecentRequests(mobile string, within time.Duration) (int, error) {
	var count int
	err := r.DB.QueryRow(`
		SELECT COUNT(*) FROM otp_requests 
		WHERE mobile = $1 AND requested_at >= NOW() - $2::interval
	`, mobile, within.String()).Scan(&count)
	return count, err
}

func (r *OTPThrottleRepo) LogRequest(mobile string) error {
	_, err := r.DB.Exec("INSERT INTO otp_requests (mobile) VALUES ($1)", mobile)
	return err
}
