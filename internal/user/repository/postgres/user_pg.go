// internal/user/repository/postgres/user_pg.go
package postgres

import (
	"database/sql"
	"errors"

	"github.com/rezalaal/coral/internal/user/models"
)

type UserPG struct {
	DB *sql.DB
}

func NewUserPG(db *sql.DB) *UserPG {
	return &UserPG{DB: db}
}

func (r *UserPG) Create(user *models.User) error {
	query := `INSERT INTO users (name, mobile, password_hash) 
			VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	return r.DB.QueryRow(query, user.Name, user.Mobile, user.PasswordHash).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
}

func (r *UserPG) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, name, mobile, password_hash, created_at, updated_at 
	          FROM users WHERE id=$1`
	row := r.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Mobile, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPG) GetByMobile(mobile string) (*models.User, error) {
	query := `SELECT id, name, mobile, password_hash, created_at, updated_at 
	          FROM users WHERE mobile=$1`
	row := r.DB.QueryRow(query, mobile)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Mobile, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserPG) List() ([]*models.User, error) {
	query := `SELECT id, name, mobile, created_at, updated_at FROM users`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Mobile, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *UserPG) Delete(id int64) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	return err
}
