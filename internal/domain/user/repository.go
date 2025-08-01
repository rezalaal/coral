package user

type Repository interface {
	Create(u *User) error
	FindByEmail(email string) (*User, error)
	FindByMobile(mobile string) (*User, error) // جدید
}

