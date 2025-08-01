package otp

type Service interface {
	GenerateCode() string
	IsValid(code string) bool
}
