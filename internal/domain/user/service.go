package user

type Service interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (*User, error)
}

type TokenGenerator interface {
	Generate(userID int64) (string, error)
}