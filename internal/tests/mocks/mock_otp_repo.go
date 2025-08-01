package mocks

import "errors"

type InMemoryOTPRepo struct {
	data map[string]struct {
		Code      string
		ExpiresAt int64
	}
}

func NewInMemoryOTPRepo() *InMemoryOTPRepo {
	return &InMemoryOTPRepo{data: make(map[string]struct {
		Code      string
		ExpiresAt int64
	})}
}

func (r *InMemoryOTPRepo) Save(mobile, code string, expiresAt int64) error {
	r.data[mobile] = struct {
		Code      string
		ExpiresAt int64
	}{code, expiresAt}
	return nil
}

func (r *InMemoryOTPRepo) Get(mobile string) (string, int64, error) {
	val, ok := r.data[mobile]
	if !ok {
		return "", 0, errors.New("not found")
	}
	return val.Code, val.ExpiresAt, nil
}
