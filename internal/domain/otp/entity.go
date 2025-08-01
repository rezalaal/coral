package otp

type OTP struct {
	Mobile  string
	Code    string
	Expires int64
}

type Repository interface {
	Save(mobile string, code string, expires int64) error
	Get(mobile string) (code string, expires int64, err error) 
	Find(mobile string) (*OTP, error)
}


