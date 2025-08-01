package otp

import "time"

type ThrottleRepository interface {
	CountRecentRequests(mobile string, within time.Duration) (int, error)
	LogRequest(mobile string) error
}
