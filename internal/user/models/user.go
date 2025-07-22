// internal/models/user.go
package models

import "time"

type User struct {
	ID           int64     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Mobile       string    `db:"mobile" json:"mobile"`
	PasswordHash string    `db:"password_hash,omitempty" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
