// internal/repository/user_repository.go
package repository

import "github.com/rezalaal/coral/internal/user/models"

// UserRepository defines methods to interact with users in the database.
type UserRepository interface {
    Create(user *models.User) error
    GetByID(id int) (*models.User, error)
    GetByEmail(email string) (*models.User, error)
    List() ([]*models.User, error)
    Delete(id int) error
}
